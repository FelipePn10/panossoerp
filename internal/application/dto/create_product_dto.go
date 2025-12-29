package dto

import "github.com/google/uuid"

type CreateProductDTO struct {
	Code      string
	GroupCode string
	Name      string
	CreatedBy uuid.UUID
}
