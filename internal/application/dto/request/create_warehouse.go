package request

import (
	"github.com/google/uuid"
)

type CreateWarehouseRequestDTO struct {
	Name        string
	Description string
	Code        string
	Type        string
	CreatedBy   uuid.UUID
}
