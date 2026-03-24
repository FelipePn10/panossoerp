package usecase

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/application/ports"
	errorsuc "github.com/FelipePn10/panossoerp/internal/application/usecase/errors"
	"github.com/FelipePn10/panossoerp/internal/domain/items/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/items/repository"
)

type CreateItemUseCase struct {
	repo repository.ItemRepository
	auth ports.AuthService
}

func (uc *CreateItemUseCase) Execute(
	ctx context.Context,
	item *entity.Item,
) (*entity.Item, error) {
	if !uc.auth.CanCreateItem(ctx) {
		return nil, errorsuc.ErrUnauthorized
	}

	created, err := uc.repo.Create(ctx, item)
	if err != nil {
		return nil, err
	}

	return created, nil
}
