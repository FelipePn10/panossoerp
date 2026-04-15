package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/FelipePn10/panossoerp/internal/application/dto/request"
	"github.com/FelipePn10/panossoerp/internal/application/ports"
	errorsuc "github.com/FelipePn10/panossoerp/internal/application/usecase/errors"
	"github.com/FelipePn10/panossoerp/internal/domain/structure/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/structure/repository"
)

// CreateStructureComponentUseCase adiciona um componente (filho) à estrutura
type CreateStructureComponentUseCase struct {
	repo repository.ItemStructureRepository
	auth ports.AuthService
}

func (uc *CreateStructureComponentUseCase) Execute(
	ctx context.Context,
	dto request.CreateStructureComponentDTO,
) (*entity.ItemStructure, error) {
	if !uc.auth.CanCreateStructure(ctx) {
		return nil, errorsuc.ErrUnauthorized
	}

	parentExists, err := uc.repo.ItemExists(ctx, dto.ParentItemID)
	if err != nil {
		return nil, fmt.Errorf("checking parent item: %w", err)
	}
	if !parentExists {
		return nil, fmt.Errorf("parent item %d not found", dto.ParentItemID)
	}

	childExists, err := uc.repo.ItemExists(ctx, dto.ChildItemID)
	if err != nil {
		return nil, fmt.Errorf("checking child item: %w", err)
	}
	if !childExists {
		return nil, fmt.Errorf("child item %d not found", dto.ChildItemID)
	}

	// Validação da referência cíclica
	// Verifica se childItemID já é ancestral de parentItemID.
	// Se sim, adicionar o filho criaria um ciclo.
	hasCycle, err := uc.repo.HasCyclicReference(ctx, dto.ParentItemID, dto.ChildItemID)
	if err != nil {
		return nil, fmt.Errorf("checking cyclic reference: %w", err)
	}
	if hasCycle {
		return nil, errors.New(
			"it is not possible to add this component: create a circular reference in the BOM tree",
		)
	}

	structure, err := entity.NewItemStructure(
		dto.ParentItemID,
		dto.ChildItemID,
		dto.ParentCode,
		dto.ChildCode,
		dto.ParentMask,
		dto.Quantity,
		dto.UnitOfMeasurement,
		dto.LossPercentage,
		dto.Position,
		dto.Notes,
		dto.CreatedBy,
	)
	if err != nil {
		return nil, err
	}

	return uc.repo.Create(ctx, structure)
}
