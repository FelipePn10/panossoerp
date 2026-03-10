package entity

import (
	"time"

	"github.com/google/uuid"
)

// raw materials and/or components of the industry's main products
type Component struct {
	ID        int64
	Name      string
	GroupCode string
	Code      string
	Warehouse int64
	CreatedBy uuid.UUID
	CreatedAt time.Time
}

type ComponentMask struct {
	ID          int64
	ComponentID int64
	Mask        string
	MaskHash    string
	CreatedBy   uuid.UUID
	CreatedAt   time.Time
}
