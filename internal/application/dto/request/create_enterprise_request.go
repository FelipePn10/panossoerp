package request

import "github.com/google/uuid"

type CreateEnterpriseDTO struct {
	Code      int       `json:"code"`
	Name      string    `json:"name"`
	CreatedBy uuid.UUID `json:"created_by"`
}
