package request

import "github.com/google/uuid"

type CreateQuestionRequestDTO struct {
	Name      string    `json:"name"`
	CreatedBy uuid.UUID `json:"created_by"`
}
