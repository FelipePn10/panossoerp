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
	ID        uuid.UUID
	Code      valueobject.ProductCode
	GroupCode string
	Name      string
	CreatedBy uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ProductMask struct {
	ID         uuid.UUID
	ProductID  uuid.UUID
	Mask       string
	MaskHash   string
	BusinessID string
	CreatedBy  uuid.UUID
	CreatedAt  time.Time
}

type Component struct {
	ID        uuid.UUID
	Code      string
	Type      ContentType
	CreatedAt time.Time
}

type ComponentMask struct {
	ID          uuid.UUID
	ComponentID uuid.UUID
	Mask        string
	MaskHash    string
	BusinessID  string
	CreatedAt   time.Time
}

type MaskComposition struct {
	ParentMaskID uuid.UUID
	ChildMaskID  uuid.UUID
	Quantity     valueobject.Quantity
}

type MaterialConsumption struct {
	ComponentMaskID uuid.UUID
	MaterialID      uuid.UUID
	Quantity        valueobject.Quantity
	Unit            string
}
