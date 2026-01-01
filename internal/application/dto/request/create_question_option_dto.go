package request

import "github.com/google/uuid"

type CreateQuestionOptionRequest struct {
	Value      string    `json:"value"`
	CreatedBy  uuid.UUID `json:"createdby"`
	QuestionId int64     `json:"question_id"`
}
