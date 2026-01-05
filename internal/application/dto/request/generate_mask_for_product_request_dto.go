package request

import (
	"github.com/google/uuid"
)

type GenerateMaskForProductRequestDTO struct {
	ProductCode string
	Options     string
	CreatedBy   uuid.UUID
}
