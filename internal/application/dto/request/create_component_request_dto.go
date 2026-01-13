package request

import "github.com/google/uuid"

type CreateComponentRequestDTO struct {
	Name      string    `json:"name"`
	GroupCode string    `json:"group_code"`
	Warehouse int64     `json:"warehouse_id"`
	CreatedBy uuid.UUID `json:"created_by"`
}
