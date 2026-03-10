package usecase

import (
	"context"
	"errors"
	"strings"

	errorsuc "github.com/FelipePn10/panossoerp/internal/application/usecase/errors"
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
	name = strings.TrimSpace(name)
	code = strings.TrimSpace(code)

	if name == "" || code == "" {
		return nil, errorsuc.ErrInvalidSearchParams
	}

	product, err := uc.repo.FindByNameAndCode(ctx, name, code)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, errorsuc.ErrProductNotFound
		}
		return nil, err
	}
	return product, nil
}
