package item

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/enums/types"
	"github.com/FelipePn10/panossoerp/internal/domain/items/entity"
	"github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"
)

func (r *repositoryItemSQLC) Create(
	ctx context.Context,
	item *entity.Item,
) (*entity.Item, error) {
	params := sqlc.CreateItemParams{
		ID:          item.ID,
		WarehouseID: item.WarehouseID,
		Code:        item.Code,
		Name:        item.Name,
		Description: item.Description,
		Type:        int16(item.Type),
		Status:      int16(item.Status),
		Health:      int16(item.Health),
		CreatedBy:   item.CreatedBy,
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
		Type:        types.Type(dbItem.Type),
		Status:      types.Status(dbItem.Status),
		Health:      types.Health(dbItem.Health),
		CreatedBy:   dbItem.CreatedBy,
		CreatedAt:   dbItem.CreatedAt,
	}, nil
}
