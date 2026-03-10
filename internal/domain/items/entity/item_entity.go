package entity

import (
	"time"

	"github.com/FelipePn10/panossoerp/internal/domain/enums/types"
	"github.com/google/uuid"
)

type Item struct {
	ID          int64
	WarehouseID string
	Code        string
	Name        string
	Description string
	// Price

	Type   types.Type
	Status types.Status
	Health types.Health

	CreatedBy uuid.UUID
	CreatedAt time.Time
}
