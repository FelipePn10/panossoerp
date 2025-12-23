package dto

import "github.com/google/uuid"

type CreateProductDTO struct {
	Code      string
	GroupCode int16
	Name      string
	CreatedBy uuid.UUID
}
