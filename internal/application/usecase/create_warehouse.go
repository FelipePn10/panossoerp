package usecase

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/application/dto/request"
	"github.com/FelipePn10/panossoerp/internal/application/ports"
	errorsuc "github.com/FelipePn10/panossoerp/internal/application/usecase/errors"
	"github.com/FelipePn10/panossoerp/internal/domain/warehouse/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/warehouse/repository"
)

type CreateWarehouseUseCase struct {
	repo repository.WarehouseRepository
	auth ports.AuthService
}

func (uc *CreateWarehouseUseCase) Execute(
	ctx context.Context,
	dto request.CreateWarehouseRequestDTO,
) (*entity.Warehouse, error) {
	if !uc.auth.CanCreateWarehouse(ctx) {
		return nil, errorsuc.ErrUnauthorized
	}

	exists, err := uc.repo.ExistsWarehouseByName(ctx, dto.Name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errorsuc.ErrWarehouseAlreadyExists
	}

	warehouse, err := entity.NewWarehouse(
		dto.Name,
		dto.Description,
		dto.Code,
		dto.Type,
		dto.CreatedBy,
	)
	if err != nil {
		return nil, err
	}

	create, err := uc.repo.Create(ctx, warehouse)
	if err != nil {
		return nil, err
	}
	return create, nil
}
