package item

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/items/entity"
)

func (r *repositoryItemSQLC) Create(
	ctx context.Context,
	item *entity.Item,
) (*entity.Item, error) {
	params := sqlc.CreateItemProductsParams{
		ID:          item.ID,
		WarehouseID: item.WarehouseID,
		Code:        item.Code,
		Name:        item.Name,
		Description: item.Description,
		Type:        item.Type,
		Status:      item.Status,
		Health:      item.Health,
		CreatedBy:   item.CreatedBy,
		CreatedAt:   item.CreatedAt,
	}

	dbItem, err := r.q.CreateItem(ctx, params)
	if err != nil {
		return nil, err
	}

	return &entity.Item{
		ID:          dbItem.ID,
		WarehouseID: dbItem.WarehouseID,
		Code:        dbItem.Code,
		Name:        dbItem.Name,
		Description: dbItem.Description,
		Type:        dbItem.Type,
		Status:      dbItem.Status,
		Health:      dbItem.Health,
		CreatedBy:   dbItem.CreatedBy,
		CreatedAt:   dbItem.CreatedAt,
	}, nil
}
