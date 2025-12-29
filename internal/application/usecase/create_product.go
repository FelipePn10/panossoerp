package usecase

import (
	"context"
	"time"

	"github.com/FelipePn10/panossoerp/internal/application/dto/request"
	"github.com/FelipePn10/panossoerp/internal/domain/product/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/product/repository"
	"github.com/FelipePn10/panossoerp/internal/domain/product/valueobject"
)

type CreateProductUseCase struct {
	repo repository.ProductRepository
}

func (uc *CreateProductUseCase) Execute(
	ctx context.Context,
	dto request.CreateProductDTO,
) error {
	now := time.Now()

	code, err := valueobject.NewProductCode(dto.GroupCode, now)
	if err != nil {
		return err
	}

	product, err := entity.NewProduct(
		code,
		dto.GroupCode,
		dto.Name,
		dto.CreatedBy,
	)
	if err != nil {
		return err
	}

	return uc.repo.Save(ctx, product)

}
