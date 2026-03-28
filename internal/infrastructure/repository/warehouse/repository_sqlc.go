package warehouse

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/warehouse/entity"
	"github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"
	mapper "github.com/FelipePn10/panossoerp/internal/infrastructure/mapper/warehouse"
)

func (r *repositoryWarehouseSQLC) Create(
	ctx context.Context,
	warehouse *entity.Warehouse,
) (*entity.Warehouse, error) {
	params := sqlc.CreateWarehouseParams{
		Code:                warehouse.Code,
		Description:         warehouse.Description,
		Location:            mapper.WarehouseLocationToDB(warehouse.Location),
		Type:                mapper.WarehouseTypeToDB(warehouse.Type),
		Disposition:         warehouse.Disposition,
		ReservationsAllowed: warehouse.ReservationsAllowed,
		CreatedBy:           warehouse.CreatedBy,
	}
	dbWarehouse, err := r.q.CreateWarehouse(ctx, params)
	if err != nil {
		return nil, err
	}

	return &entity.Warehouse{
		Code:                dbWarehouse.Code,
		Description:         dbWarehouse.Description,
		Location:            mapper.WarehouseLocationToDomain(dbWarehouse.Location),
		Type:                mapper.WarehouseTypeToDomain(dbWarehouse.Type),
		Disposition:         dbWarehouse.Disposition,
		ReservationsAllowed: dbWarehouse.ReservationsAllowed,
		CreatedBy:           dbWarehouse.CreatedBy,
		CreatedAt:           dbWarehouse.CreatedAt,
	}, nil
}
