package request

import "github.com/google/uuid"

type CreateQuestionRequestDTO struct {
	Name      string
	CreatedBy uuid.UUID
}
