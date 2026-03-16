package warehouse

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/warehouse/entity"
)

func (r *repositoryWarehouseSQLC) Create(
	ctx context.Context,
	warehouse *entity.Warehouse,
) (*entity.Warehouse, error) {
	params := sqlc.CreateWarehouseParams{}

	dbWarehouse, err := r.q.CreateWarehouse(ctx, params)
	if err != nil {
		return nil, err
	}

	return &entity.Warehouse{
		Name: ,
		Description: ,
		Code: ,
		Type: ,
		CreatedBy: ,
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
