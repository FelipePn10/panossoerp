package entity

import (
	"time"

	"github.com/FelipePn10/panossoerp/internal/domain/enums/types"
	"github.com/FelipePn10/panossoerp/internal/domain/items/entity"
	"github.com/google/uuid"
)

type Warehouse struct {
	ID          int32
	Name        string
	Description string
	Code        string
	Type        types.TypeWarehouse

	Items_List []entity.Item

	CreatedBy uuid.UUID
	CreatedAt time.Time
}
