package request

import "github.com/google/uuid"

type CreateProductDTO struct {
	GroupCode string    `json:"group_code"`
	Name      string    `json:"name"`
	CreatedBy uuid.UUID `json:"created_by"`
}
