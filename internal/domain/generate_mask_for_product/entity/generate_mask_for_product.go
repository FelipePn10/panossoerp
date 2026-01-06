package entity

import (
	"time"

	"github.com/FelipePn10/panossoerp/internal/domain/generate_mask_for_product/valueobject"
	"github.com/google/uuid"
)

type ProductMask struct {
	ID          int64
	ProductCode string
	Mask        string
	MaskHash    string
	CreatedBy   uuid.UUID
	CreatedAt   time.Time
	Answers     []valueobject.MaskAnswer
}
