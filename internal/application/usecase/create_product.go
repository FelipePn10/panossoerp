package usecase

import (
	"context"
	"time"

	"github.com/FelipePn10/panossoerp/internal/application/dto/request"
	"github.com/FelipePn10/panossoerp/internal/application/ports"
	errorsuc "github.com/FelipePn10/panossoerp/internal/application/usecase/errors"
	"github.com/FelipePn10/panossoerp/internal/domain/product/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/product/repository"
	"github.com/FelipePn10/panossoerp/internal/domain/product/valueobject"
)

type CreateProductUseCase struct {
	repo repository.ProductRepository
	auth ports.AuthService
}

func (uc *CreateProductUseCase) Execute(
	ctx context.Context,
	dto request.CreateProductDTO,
) (*entity.Product, error) {
	if !uc.auth.CanCreateProduct(ctx) {
		return nil, errorsuc.ErrUnauthorized
	}

	now := time.Now()
	code, err := valueobject.NewProductCode(dto.GroupCode, now)
	if err != nil {
		return nil, err
	}

	exists, err := uc.repo.ExistsByCode(ctx, code.String())
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errorsuc.ErrProductAlreadyExists
	}

	product, err := entity.NewProduct(
		code.String(),
		dto.GroupCode,
		dto.Name,
		dto.CreatedBy,
	)
	if err != nil {
		return nil, err
	}

	saved, err := uc.repo.Save(ctx, product)
	if err != nil {
		return nil, err
	}

	return saved, nil
}
