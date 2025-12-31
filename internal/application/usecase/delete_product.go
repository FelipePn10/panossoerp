package usecase

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/product/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/product/repository"
)

type DeleteProductUseCase struct {
	repo repository.ProductRepository
}

func (uc *DeleteProductUseCase) Execute(
	ctx context.Context,
	id int64,
) error {
	if err := entity.ValidateProductDeletion(id); err != nil {
		return err
	}
	return uc.repo.Delete(ctx, id)
}
