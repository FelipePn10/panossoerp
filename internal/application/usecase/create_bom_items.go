package usecase

import (
	"context"
	"errors"

	"github.com/FelipePn10/panossoerp/internal/application/dto/request"
	"github.com/FelipePn10/panossoerp/internal/application/ports"
	errorsuc "github.com/FelipePn10/panossoerp/internal/application/usecase/errors"
	bom "github.com/FelipePn10/panossoerp/internal/domain/bom/repository"
	"github.com/FelipePn10/panossoerp/internal/domain/bom_items/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/bom_items/repository"
)

type CreateBomItemUseCase struct {
	repo repository.BomItemsRepository
	bom  bom.BomRepository
	auth ports.AuthService
}

func (uc *CreateBomItemUseCase) Execute(
	ctx context.Context,
	dto request.CreateBomItemsRequestDTO,
) (*entity.BomItems, error) {
	if !uc.auth.CanCreateBomItems(ctx) {
		return nil, errorsuc.ErrUnauthorized
	}

	bomItem, err := entity.NewBomItems(
		dto.BomID,
		dto.ComponentID,
		dto.Quantity,
		dto.Uom,
		dto.ScrapPercent,
		dto.OperationID,
		dto.MaskComponent,
	)
	if err != nil {
		if errors.Is(err, repository.ErrInvalidBomItems) {
			return nil, errorsuc.ErrCreateBomItem
		}
		return nil, err
	}
	created, err := uc.repo.Create(ctx, bomItem)
	if err != nil {
		return nil, err
	}

	return created, nil
}
