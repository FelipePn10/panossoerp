package usecase

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/application/dto/request"
	"github.com/FelipePn10/panossoerp/internal/application/ports"
	errorsuc "github.com/FelipePn10/panossoerp/internal/application/usecase/errors"
	"github.com/FelipePn10/panossoerp/internal/domain/items/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/items/repository"
	"github.com/FelipePn10/panossoerp/internal/domain/items/valueobject"
)

type FindItemByCode struct {
	repo repository.ItemRepository
	auth ports.AuthService
}

func (uc *FindItemByCode) Execute(
	ctx context.Context,
	dto request.FindItemByCodeDTO,
) (*entity.Item, error) {
	if !uc.auth.FindItemByCode(ctx) {
		return nil, errorsuc.ErrUnauthorized
	}

	code, err := valueobject.NewItemCode(int64(dto.Code))
	if err != nil {
		return nil, err
	}

	item, err := uc.repo.FindItemByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	return item, nil
}
