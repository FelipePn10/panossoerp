package usecase

import (
	"context"
	"errors"

	"github.com/FelipePn10/panossoerp/internal/domain/product/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/product/repository"
)

type FindProductByNameAndCode struct {
	repo repository.ProductRepository
}

func (uc *FindProductByNameAndCode) Execute(
	ctx context.Context,
	name string,
	code string,
) (*entity.Product, error) {
	if name == "" || code == "" {
		return nil, errors.New("name and code is required")
	}

	return uc.repo.FindByNameAndCode(ctx, name, code)
}
