package dto

import "github.com/google/uuid"

type CreateProductDTO struct {
	GroupCode string
	Name      string
	CreatedBy uuid.UUID
}
