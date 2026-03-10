package usecase

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/application/ports"
	errorsuc "github.com/FelipePn10/panossoerp/internal/application/usecase/errors"
	"github.com/FelipePn10/panossoerp/internal/domain/product/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/product/repository"
)

type DeleteProductUseCase struct {
	repo repository.ProductRepository
	auth ports.AuthService
}

func (uc *DeleteProductUseCase) Execute(
	ctx context.Context,
	id int64,
) error {
	if !uc.auth.CanDeleteProduct(ctx) {
		return errorsuc.ErrUnauthorized
	}

	if err := entity.ValidateProductDeletion(id); err != nil {
		return err
	}
	return uc.repo.Delete(ctx, id)
}
