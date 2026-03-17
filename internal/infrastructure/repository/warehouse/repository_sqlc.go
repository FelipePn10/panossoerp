package warehouse

import (
	"context"
	"database/sql"
	"errors"

	"github.com/FelipePn10/panossoerp/internal/domain/warehouse/entity"
	"github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"
)

func (r *repositoryWarehouseSQLC) Create(
	ctx context.Context,
	warehouse *entity.Warehouse,
) (*entity.Warehouse, error) {
	params := sqlc.CreateWarehouseParams{
		Name:        warehouse.Name,
		Description: warehouse.Description,
		Code:        warehouse.Code,
		Types:       warehouse.Type,
		CreatedBy:   warehouse.CreatedBy,
	}

	dbWarehouse, err := r.q.CreateWarehouse(ctx, params)
	if err != nil {
		return nil, err
	}

	return &entity.Warehouse{
		ID:          int32(dbWarehouse.ID),
		Name:        dbWarehouse.Name,
		Description: dbWarehouse.Description,
		Code:        dbWarehouse.Code,
		Type:        dbWarehouse.Types,
		CreatedBy:   dbWarehouse.CreatedBy,
	}, nil
}

func (r *repositoryWarehouseSQLC) ExistsWarehouseByName(
	ctx context.Context,
	name string,
) (bool, error) {
	_, err := r.q.ExistsWarehouseByName(ctx, name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, err
		}
		return false, err
	}
	return true, nil
}
