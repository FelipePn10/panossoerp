package repository

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/items/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/items/valueobject"
)

type ItemRepository interface {
	Create(ctx context.Context, item *entity.Item) (*entity.Item, error)
	FindItemByCode(ctx context.Context, code valueobject.ItemCode) (*entity.Item, error)
}
