package repository

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/items/entity"
)

type ItemRepository interface {
	Create(ctx context.Context, item *entity.Item) (*entity.Item, error)
	ExistsItemByCode(ctx context.Context, code int64)
}
