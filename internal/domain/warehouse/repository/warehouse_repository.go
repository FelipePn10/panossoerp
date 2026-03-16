package repository

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/warehouse/entity"
)

type WarehouseRepository interface {
	Create(ctx context.Context, warehouse *entity.Warehouse) (*entity.Warehouse, error)
	ExistsWarehouseByName(ctx context.Context, name string) (bool, error)
}
