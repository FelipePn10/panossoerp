package entity

import (
	"time"

	"github.com/google/uuid"
)

type ProductMask struct {
	ID          int64
	ProductCode string
	Mask        string
	MaskHash    string
	CreatedBy   uuid.UUID
	CreatedAt   time.Time
}
