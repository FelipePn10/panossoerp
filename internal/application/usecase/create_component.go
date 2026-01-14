package usecase

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/application/dto/request"
	"github.com/FelipePn10/panossoerp/internal/application/ports"
	errorsuc "github.com/FelipePn10/panossoerp/internal/application/usecase/errors"
	"github.com/FelipePn10/panossoerp/internal/domain/component/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/component/repository"
	"github.com/FelipePn10/panossoerp/internal/domain/component/valueobject"
)

type CreateComponentUseCase struct {
	repo repository.ComponentRepository
	auth ports.AuthService
}

func (uc *CreateComponentUseCase) Execute(
	ctx context.Context,
	dto request.CreateComponentRequestDTO,
) (*entity.Component, error) {

	if !uc.auth.CanCreateComponent(ctx) {
		return nil, errorsuc.ErrUnauthorized
	}

	code, err := valueobject.NewComponentCode(dto.GroupCode)
	if err != nil {
		return &entity.Component{}, err
	}

	exists, err := uc.repo.ExistsByCode(ctx, code.String())
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errorsuc.ErrComponentAlreadyExists
	}

	// warehouse, err := uc.warehouseRepo.FindByID(ctx, dto.WarehouseID)
	// if err != nil || warehouse == nil {
	// 	return nil, ErrWarehouseNotFound
	// }

	component, err := entity.NewComponent(
		code.String(),
		dto.GroupCode,
		dto.Name,
		dto.Warehouse,
		dto.CreatedBy,
	)
	if err != nil {
		return nil, err
	}

	saved, err := uc.repo.Save(ctx, component)
	if err != nil {
		return nil, err
	}
	return saved, nil
}
