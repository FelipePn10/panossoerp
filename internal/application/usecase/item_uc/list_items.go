package item_uc

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/application/ports"
	errorsuc "github.com/FelipePn10/panossoerp/internal/application/usecase/errors"
	"github.com/FelipePn10/panossoerp/internal/domain/items/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/items/repository"
)

type ListItemsUseCase struct {
	Repo repository.ItemRepository
	Auth ports.AuthService
}

func NewListItemsUseCase(repo repository.ItemRepository, auth ports.AuthService) *ListItemsUseCase {
	return &ListItemsUseCase{Repo: repo, Auth: auth}
}

func (uc *ListItemsUseCase) Execute(ctx context.Context) ([]*entity.Item, error) {
	if !uc.Auth.FindItemByCode(ctx) {
		return nil, errorsuc.ErrUnauthorized
	}
	return uc.Repo.ListAll(ctx)
}
