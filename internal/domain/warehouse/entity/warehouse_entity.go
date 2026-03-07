package entity

import (
	"time"

	"github.com/FelipePn10/panossoerp/internal/domain/items/entity"
	"github.com/google/uuid"
)

type Warehouse struct {
	ID   int64
	Name string

	Items_List []entity.Item

	CreatedBy uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}
