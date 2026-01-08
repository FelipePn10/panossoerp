package entity

import (
	"time"

	"github.com/FelipePn10/panossoerp/internal/domain/product/valueobject"
	"github.com/google/uuid"
)

type ContentType string

const (
	ContentStructure ContentType = "STRUCTURE"
	ContentSet       ContentType = "SET"
	ContentItem      ContentType = "ITEM"
)

type Product struct {
	ID        int64
	Code      string
	GroupCode string
	Name      string
	UOM       string
	CreatedBy uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ProductMask struct {
	ID        int64
	ProductID int64
	Mask      string
	MaskHash  string
	CreatedBy uuid.UUID
	CreatedAt time.Time
}

type Component struct {
	ID        int64
	Code      string
	Type      ContentType
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

type MaskComposition struct {
	ParentMaskID int64
	ChildMaskID  int64
	Quantity     valueobject.Quantity
}

type MaterialConsumption struct {
	ComponentMaskID int64
	MaterialID      int64
	Quantity        valueobject.Quantity
	Unit            string
}
