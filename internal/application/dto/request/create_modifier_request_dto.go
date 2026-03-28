package request

import "github.com/google/uuid"

type CreateModifierDTO struct {
	Description string    `json:"description"`
	CreatedBy   uuid.UUID `json:"created_by"`
}
