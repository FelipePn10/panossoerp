package request

import (
	"github.com/FelipePn10/panossoerp/internal/domain/enums/types"
	"github.com/google/uuid"
)

type CreateItemDTO struct {
	WarehouseID int64        `json:"warehouse_id"`
	Name        string       `json:"name"`
	Desc        string       `json:"desc"`
	Type        types.Type   `json:"type"`
	Status      types.Status `json:"status"`
	Health      types.Health `json:"health"`
	CreatedBy   uuid.UUID    `json:"created_by"`
}
