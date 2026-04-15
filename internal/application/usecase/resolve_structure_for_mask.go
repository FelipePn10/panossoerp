package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/FelipePn10/panossoerp/internal/application/dto/request"
	"github.com/FelipePn10/panossoerp/internal/application/dto/response"
	"github.com/FelipePn10/panossoerp/internal/application/ports"
	errorsuc "github.com/FelipePn10/panossoerp/internal/application/usecase/errors"
	"github.com/FelipePn10/panossoerp/internal/domain/structure/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/structure/repository"
	"github.com/FelipePn10/panossoerp/internal/domain/structure/valueobject"
	mapper "github.com/FelipePn10/panossoerp/internal/infrastructure/mapper/structure"
)

// ResolveStructureForMaskUseCase é o caso de uso central da estrutura de produtos.
//
// Dado um item raiz e sua máscara (ex.: "100#100#50"), resolve a árvore BOM
// completa aplicando as seguintes regras em cada nível:
//
//  1. PRIORIDADE: componente específico (parent_mask = máscara calculada)
//     sobrepõe o componente genérico (parent_mask IS NULL) para o mesmo
//     item filho.
//
//  2. PROPAGAÇÃO DE MÁSCARA: as respostas da máscara do pai são repassadas
//     para os filhos que compartilham as mesmas perguntas.
//     Ex.: pai tem Q1=100, Q2=100, Q3=50; filho tem Q1 e Q2 →
//     máscara do filho = "100#100"
//
//  3. FALLBACK GENÉRICO: se o filho não tem perguntas configuradas ou
//     nenhuma máscara foi calculada, usa-se a versão genérica.
type ResolveStructureForMaskUseCase struct {
	repo repository.ItemStructureRepository
	auth ports.AuthService
}

func (uc *ResolveStructureForMaskUseCase) Execute(
	ctx context.Context,
	dto request.ResolveStructureForMaskDTO,
) (*response.StructureTreeResponse, error) {
	if !uc.auth.ResolveStructureForMask(ctx) {
		return nil, errorsuc.ErrUnauthorized
	}

	if dto.RootMaskValue == "" {
		return nil, errors.New("root mask value – required for structure resolution")
	}

	rootExists, err := uc.repo.ItemExists(ctx, dto.RootItemID)
	if err != nil {
		return nil, fmt.Errorf("checking root item: %w", err)
	}
	if !rootExists {
		return nil, fmt.Errorf("root %d not found", dto.RootItemID)
	}

	rootCode, rootDesc, err := uc.repo.GetItemCodeAndDesc(ctx, dto.RootItemID)
	if err != nil {
		return nil, fmt.Errorf("retrieving data from the root item: %w", err)
	}

	rootAnswers, err := uc.repo.GetMaskAnswersByItemAndValue(ctx, dto.RootItemID, dto.RootMaskValue)
	if err != nil {
		return nil, fmt.Errorf("seeking answers from the root mask: %w", err)
	}
	if len(rootAnswers) == 0 {
		return nil, fmt.Errorf(
			"mask '%s' not found for item %d; generate the mask before resolving the structure",
			dto.RootMaskValue, dto.RootItemID,
		)
	}

	// resolve árvore recursivamente
	visited := make(map[int64]bool)
	nodes, err := uc.resolveTree(ctx, dto.RootItemID, dto.RootMaskValue, rootAnswers, 1, visited)
	if err != nil {
		return nil, err
	}

	responseNodes := make([]*response.StructureTreeNodeResponse, 0, len(nodes))
	for _, n := range nodes {
		responseNodes = append(responseNodes, mapper.MapNodeToResponse(n))
	}

	rootMask := dto.RootMaskValue
	return &response.StructureTreeResponse{
		RootItemID:  dto.RootItemID,
		RootCode:    rootCode,
		RootDesc:    rootDesc,
		RootMask:    &rootMask,
		Components:  responseNodes,
		TotalLevels: mapper.MaxLevel(responseNodes) + 1,
		TotalNodes:  mapper.CountNodes(responseNodes),
	}, nil
}

// resolveTree resolve recursivamente os filhos de um nó para uma máscara específica.
//
// parentAnswers são as respostas já carregadas para o pai atual; são usadas
// para propagar a máscara para os filhos.
func (uc *ResolveStructureForMaskUseCase) resolveTree(
	ctx context.Context,
	parentItemID int64,
	parentMaskValue string,
	parentAnswers []valueobject.MaskAnswer,
	level int,
	visited map[int64]bool,
) ([]*valueobject.StructureNode, error) {
	if level > maxBOMDepth {
		return nil, fmt.Errorf("maximum depth of BOM (%d) reached", maxBOMDepth)
	}
	if visited[parentItemID] {
		return nil, nil
	}
	visited[parentItemID] = true
	defer func() { delete(visited, parentItemID) }()

	// Busca filhos: genéricos + específicos para a máscara atual
	allChildren, err := uc.repo.GetDirectChildrenForMask(ctx, parentItemID, parentMaskValue)
	if err != nil {
		return nil, fmt.Errorf("searching for children of item %d (mask '%s'): %w", parentItemID, parentMaskValue, err)
	}

	// Deduplica: para cada child_item_id, o componente ESPECÍFICO tem prioridade
	// sobre o genérico.
	// GetDirectChildrenForMask já retorna específicos primeiro (ORDER BY).
	resolved := deduplicateByChildID(allChildren)

	nodes := make([]*valueobject.StructureNode, 0, len(resolved))

	for _, comp := range resolved {
		code, desc, err := uc.repo.GetItemCodeAndDesc(ctx, comp.ChildItemID)
		if err != nil {
			return nil, fmt.Errorf("searching for item data %d: %w", comp.ChildItemID, err)
		}

		// Propaga a máscara do pai para o filho
		childMask, childAnswers, err := uc.computeChildMask(ctx, comp.ChildItemID, parentAnswers)
		if err != nil {
			return nil, err
		}

		node := valueobject.NewStructureNode(comp, code, desc, level, childMask)

		// Recursão
		var children []*valueobject.StructureNode
		if childMask != nil && len(childAnswers) > 0 {
			// Filho tem máscara calculada: resolve a sub-árvore com ela
			children, err = uc.resolveTree(ctx, comp.ChildItemID, *childMask, childAnswers, level+1, visited)
		} else {
			// Filho não tem máscara calculável: usa somente filhos genéricos
			children, err = uc.resolveTreeGeneric(ctx, comp.ChildItemID, level+1, visited)
		}
		if err != nil {
			return nil, err
		}
		for _, child := range children {
			node.AddChild(child)
		}

		nodes = append(nodes, node)
	}

	return nodes, nil
}

// computeChildMask calcula a máscara do filho com base nas respostas do pai.
// Retorna (nil, nil, nil) quando o filho não tem perguntas configuradas.
func (uc *ResolveStructureForMaskUseCase) computeChildMask(
	ctx context.Context,
	childItemID int64,
	parentAnswers []valueobject.MaskAnswer,
) (*string, []valueobject.MaskAnswer, error) {
	childQuestions, err := uc.repo.GetItemQuestions(ctx, childItemID)
	if err != nil {
		return nil, nil, fmt.Errorf("searching for questions from the item %d: %w", childItemID, err)
	}
	if len(childQuestions) == 0 {
		return nil, nil, nil
	}

	// Propaga as respostas do pai para o filho
	childMask := valueobject.PropagateMask(parentAnswers, childQuestions)
	if childMask == nil {
		return nil, nil, nil
	}

	// Busca as respostas da máscara do filho para propagar ao próximo nível
	childAnswers, err := uc.repo.GetMaskAnswersByItemAndValue(ctx, childItemID, *childMask)
	if err != nil {
		// Máscara calculada mas não existe no BD → fallback para genérico
		return nil, nil, nil
	}

	return childMask, childAnswers, nil
}

// resolveTreeGeneric resolve a sub-árvore de um nó usando apenas filhos genéricos.
// Usado como fallback quando a máscara não pode ser propagada.
func (uc *ResolveStructureForMaskUseCase) resolveTreeGeneric(
	ctx context.Context,
	parentItemID int64,
	level int,
	visited map[int64]bool,
) ([]*valueobject.StructureNode, error) {
	if level > maxBOMDepth || visited[parentItemID] {
		return nil, nil
	}
	visited[parentItemID] = true
	defer func() { delete(visited, parentItemID) }()

	children, err := uc.repo.GetGenericChildren(ctx, parentItemID)
	if err != nil {
		return nil, fmt.Errorf("searching for generic children of the item %d: %w", parentItemID, err)
	}

	nodes := make([]*valueobject.StructureNode, 0, len(children))
	for _, comp := range children {
		code, desc, err := uc.repo.GetItemCodeAndDesc(ctx, comp.ChildItemID)
		if err != nil {
			return nil, fmt.Errorf("searching for item data %d: %w", comp.ChildItemID, err)
		}
		node := valueobject.NewStructureNode(comp, code, desc, level, nil)

		subChildren, err := uc.resolveTreeGeneric(ctx, comp.ChildItemID, level+1, visited)
		if err != nil {
			return nil, err
		}
		for _, sc := range subChildren {
			node.AddChild(sc)
		}
		nodes = append(nodes, node)
	}
	return nodes, nil
}

// deduplicateByChildID garante que para cada child_item_id só exista UM
// componente: o específico (mascarado) tem prioridade sobre o genérico.
//
// A query SQL já ordena específicos primeiro, então basta pegar o primeiro
// occurrence de cada child_item_id.
func deduplicateByChildID(components []*entity.ItemStructure) []*entity.ItemStructure {
	seen := make(map[int64]bool, len(components))
	result := make([]*entity.ItemStructure, 0, len(components))
	for _, c := range components {
		if !seen[c.ChildItemID] {
			seen[c.ChildItemID] = true
			result = append(result, c)
		}
	}
	return result
}
