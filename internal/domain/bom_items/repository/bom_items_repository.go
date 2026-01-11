package repository

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/bom_items/entity"
)

type BomItemsRepository interface {
	Create(ctx context.Context, bomitems *entity.BomItems) (*entity.BomItems, error)
}
