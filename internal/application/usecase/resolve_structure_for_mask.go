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
		return nil, errors.New("root mask value is required")
	}

	rootExists, err := uc.repo.ItemExists(ctx, dto.RootItemCode)
	if err != nil {
		return nil, fmt.Errorf("checking root item: %w", err)
	}
	if !rootExists {
		return nil, fmt.Errorf("root %d not found", dto.RootItemCode)
	}

	rootCode, rootDesc, err := uc.repo.GetItemCodeAndDesc(ctx, dto.RootItemCode)
	if err != nil {
		return nil, fmt.Errorf("retrieving root item: %w", err)
	}

	rootAnswers, err := uc.repo.GetMaskAnswersByItemAndValue(ctx, dto.RootItemCode, dto.RootMaskValue)
	if err != nil {
		return nil, fmt.Errorf("retrieving root mask answers: %w", err)
	}
	if len(rootAnswers) == 0 {
		return nil, fmt.Errorf("mask '%s' not found for item %d", dto.RootMaskValue, dto.RootItemCode)
	}

	visited := make(map[int64]bool)

	nodes, err := uc.resolveTree(
		ctx,
		dto.RootItemCode,
		dto.RootMaskValue,
		rootAnswers,
		1,
		visited,
	)
	if err != nil {
		return nil, err
	}

	responseNodes := make([]*response.StructureTreeNodeResponse, 0, len(nodes))
	for _, n := range nodes {
		responseNodes = append(responseNodes, mapper.MapNodeToResponse(n))
	}

	rootMask := dto.RootMaskValue

	return &response.StructureTreeResponse{
		RootItemCode: rootCode,
		RootCode:     rootCode,
		RootDesc:     rootDesc,
		RootMask:     &rootMask,
		Components:   responseNodes,
		TotalLevels:  mapper.MaxLevel(responseNodes) + 1,
		TotalNodes:   mapper.CountNodes(responseNodes),
	}, nil
}

func (uc *ResolveStructureForMaskUseCase) resolveTree(
	ctx context.Context,
	parentCode int64,
	parentMask string,
	parentAnswers []valueobject.MaskAnswer,
	level int,
	visited map[int64]bool,
) ([]*valueobject.StructureNode, error) {

	if level > maxBOMDepth {
		return nil, fmt.Errorf("max depth reached")
	}

	if visited[parentCode] {
		return nil, nil
	}

	visited[parentCode] = true
	defer delete(visited, parentCode)

	children, err := uc.repo.GetDirectChildrenForMask(ctx, parentCode, parentMask)
	if err != nil {
		return nil, err
	}

	resolved := deduplicateByChildID(children)

	nodes := make([]*valueobject.StructureNode, 0, len(resolved))

	for _, comp := range resolved {

		code, desc, err := uc.repo.GetItemCodeAndDesc(ctx, comp.ChildCode)
		if err != nil {
			return nil, err
		}

		childMask, childAnswers, err := uc.computeChildMask(
			ctx,
			comp.ChildCode,
			parentAnswers,
		)
		if err != nil {
			return nil, err
		}

		node := valueobject.NewStructureNode(
			comp,
			code,
			desc,
			level,
			childMask,
		)

		var sub []*valueobject.StructureNode

		if childMask != nil && len(childAnswers) > 0 {
			sub, err = uc.resolveTree(ctx, comp.ChildCode, *childMask, childAnswers, level+1, visited)
		} else {
			sub, err = uc.resolveTreeGeneric(ctx, comp.ChildCode, level+1, visited)
		}

		if err != nil {
			return nil, err
		}

		for _, s := range sub {
			node.AddChild(s)
		}

		nodes = append(nodes, node)
	}

	return nodes, nil
}

func (uc *ResolveStructureForMaskUseCase) computeChildMask(
	ctx context.Context,
	childCode int64,
	parentAnswers []valueobject.MaskAnswer,
) (*string, []valueobject.MaskAnswer, error) {

	questions, err := uc.repo.GetItemQuestions(ctx, childCode)
	if err != nil {
		return nil, nil, err
	}

	if len(questions) == 0 {
		return nil, nil, nil
	}

	mask := valueobject.PropagateMask(parentAnswers, questions)
	if mask == nil {
		return nil, nil, nil
	}

	answers, err := uc.repo.GetMaskAnswersByItemAndValue(ctx, childCode, *mask)
	if err != nil {
		return nil, nil, nil
	}

	return mask, answers, nil
}

func (uc *ResolveStructureForMaskUseCase) resolveTreeGeneric(
	ctx context.Context,
	parentCode int64,
	level int,
	visited map[int64]bool,
) ([]*valueobject.StructureNode, error) {

	if level > maxBOMDepth || visited[parentCode] {
		return nil, nil
	}

	visited[parentCode] = true
	defer delete(visited, parentCode)

	children, err := uc.repo.GetGenericChildren(ctx, parentCode)
	if err != nil {
		return nil, err
	}

	nodes := make([]*valueobject.StructureNode, 0, len(children))

	for _, comp := range children {

		code, desc, err := uc.repo.GetItemCodeAndDesc(ctx, comp.ChildCode)
		if err != nil {
			return nil, err
		}

		node := valueobject.NewStructureNode(comp, code, desc, level, nil)

		sub, err := uc.resolveTreeGeneric(ctx, comp.ChildCode, level+1, visited)
		if err != nil {
			return nil, err
		}

		for _, s := range sub {
			node.AddChild(s)
		}

		nodes = append(nodes, node)
	}

	return nodes, nil
}

func deduplicateByChildID(components []*entity.ItemStructure) []*entity.ItemStructure {
	seen := make(map[int64]bool)
	result := make([]*entity.ItemStructure, 0)

	for _, c := range components {
		if !seen[c.ChildCode] {
			seen[c.ChildCode] = true
			result = append(result, c)
		}
	}

	return result
}
