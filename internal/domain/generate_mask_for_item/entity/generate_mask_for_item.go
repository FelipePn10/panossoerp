package entity

import (
	"time"

	"github.com/FelipePn10/panossoerp/internal/domain/generate_mask_for_item/valueobject"
	"github.com/google/uuid"
)

type ItemMask struct {
	ID        int64
	ItemCode  string
	Mask      string
	MaskHash  string
	CreatedBy uuid.UUID
	CreatedAt time.Time
	Answers   []valueobject.MaskAnswer
}
