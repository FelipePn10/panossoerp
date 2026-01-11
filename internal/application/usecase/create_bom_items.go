package usecase

import (
	"context"
	"errors"

	"github.com/FelipePn10/panossoerp/internal/application/dto/request"
	errorsuc "github.com/FelipePn10/panossoerp/internal/application/usecase/errors"
	"github.com/FelipePn10/panossoerp/internal/domain/bom_items/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/bom_items/repository"
)

type CreateBomItemUseCase struct {
	repo repository.BomItemsRepository
}

func (uc *CreateBomItemUseCase) Execute(
	ctx context.Context,
	dto request.CreateBomItemsRequestDTO,
) (*entity.BomItems, error) {
	bomitems, err := entity.NewBomItems(
		dto.BomID,
		dto.ComponentID,
		dto.Quantity,
		dto.Uom,
		dto.ScrapPercent,
		dto.OperationID,
	)
	if err != nil {
		if errors.Is(err, repository.ErrInvalidBomItems) {
			return nil, errorsuc.ErrCreateBomItem
		}
		return nil, err
	}
	return uc.repo.Create(ctx, bomitems)
}
