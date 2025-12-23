package usecase

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/application/dto"
	"github.com/FelipePn10/panossoerp/internal/domain/product/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/product/repository"
	"github.com/google/uuid"
)

type CreateProductUseCase struct {
	repo repository.ProductRepository
}

func (uc *CreateProductUseCase) Execute(
	ctx context.Context,
	dto dto.CreateProductDTO,
) error {
	id := uuid.New()
	product, err := entity.NewProduct(
		id,
		dto.Code,
		dto.GroupCode,
		dto.Name,
		dto.CreatedBy,
	)
	if err != nil {
		return err
	}

	return uc.repo.Save(ctx, product)
}
