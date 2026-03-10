package request

import (
	"github.com/FelipePn10/panossoerp/internal/domain/enums/types"
	"github.com/google/uuid"
)

type CreateItemDTO struct {
	WarehouseID int64        `json:"warehouse_id"`
	Code        string       `json:"code"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Type        types.Type   `json:"type"`
	Status      types.Status `json:"status"`
	Health      types.Health `json:"health"`
	CreatedBy   uuid.UUID    `json:"created_by"`
}
