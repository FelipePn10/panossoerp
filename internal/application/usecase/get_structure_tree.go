package usecase

import (
	"context"
	"fmt"

	"github.com/FelipePn10/panossoerp/internal/application/dto/request"
	"github.com/FelipePn10/panossoerp/internal/application/dto/response"
	"github.com/FelipePn10/panossoerp/internal/application/ports"
	errorsuc "github.com/FelipePn10/panossoerp/internal/application/usecase/errors"
	"github.com/FelipePn10/panossoerp/internal/domain/structure/repository"
	"github.com/FelipePn10/panossoerp/internal/domain/structure/valueobject"
	mapper "github.com/FelipePn10/panossoerp/internal/infrastructure/mapper/structure"
)

const maxBOMDepth = 30 // Profundidade máxima para evitar loops infinitos

// GetStructureTreeUseCase retorna a árvore BOM GENÉRICA de um item raiz.
// "Genérica" significa que só são incluídos os componentes sem máscara
// (parent_mask IS NULL).
type GetStructureTreeUseCase struct {
	repo repository.ItemStructureRepository
	auth ports.AuthService
}

func (uc *GetStructureTreeUseCase) Execute(
	ctx context.Context,
	dto request.GetStructureTreeDTO,
) (*response.StructureTreeResponse, error) {
	if !uc.auth.GetStructureTree(ctx) {
		return nil, errorsuc.ErrUnauthorized
	}

	rootExists, err := uc.repo.ItemExists(ctx, dto.RootItemID)
	if err != nil {
		return nil, fmt.Errorf("checking root item: %w", err)
	}
	if !rootExists {
		return nil, fmt.Errorf("item root %d not found", dto.RootItemID)
	}

	rootCode, rootDesc, err := uc.repo.GetItemCodeAndDesc(ctx, dto.RootItemID)
	if err != nil {
		return nil, fmt.Errorf("Retrieving data from the root item: %w", err)
	}

	// monta árvore recursivamente
	// visited previne loops infinitos causados por dados inconsistentes no BD
	visited := make(map[int64]bool)
	nodes, err := uc.buildTree(ctx, dto.RootItemID, 1, visited)
	if err != nil {
		return nil, err
	}

	responseNodes := make([]*response.StructureTreeNodeResponse, 0, len(nodes))
	for _, n := range nodes {
		responseNodes = append(responseNodes, mapper.MapNodeToResponse(n))
	}

	return &response.StructureTreeResponse{
		RootItemID:  dto.RootItemID,
		RootCode:    rootCode,
		RootDesc:    rootDesc,
		RootMask:    nil, // árvore genérica nn tem máscara
		Components:  responseNodes,
		TotalLevels: mapper.MaxLevel(responseNodes) + 1,
		TotalNodes:  mapper.CountNodes(responseNodes),
	}, nil
}

// buildTree busca recursivamente os filhos genéricos de um nó.
func (uc *GetStructureTreeUseCase) buildTree(
	ctx context.Context,
	parentItemID int64,
	level int,
	visited map[int64]bool,
) ([]*valueobject.StructureNode, error) {
	if level > maxBOMDepth {
		return nil, fmt.Errorf("maximum depth of the BOM (%d levels) reached; check for cycles in the data", maxBOMDepth)
	}
	if visited[parentItemID] {
		return nil, nil
	}
	visited[parentItemID] = true
	defer func() { delete(visited, parentItemID) }() // permite reuso em ramos distintos

	components, err := uc.repo.GetGenericChildren(ctx, parentItemID)
	if err != nil {
		return nil, fmt.Errorf("searching for children of the item %d: %w", parentItemID, err)
	}

	nodes := make([]*valueobject.StructureNode, 0, len(components))
	for _, comp := range components {
		code, desc, err := uc.repo.GetItemCodeAndDesc(ctx, comp.ChildItemID)
		if err != nil {
			return nil, fmt.Errorf("searching for item data %d: %w", comp.ChildItemID, err)
		}

		node := valueobject.NewStructureNode(comp, code, desc, level, nil)

		// recursão para os filhos deste nó
		children, err := uc.buildTree(ctx, comp.ChildItemID, level+1, visited)
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
