package request

import "github.com/google/uuid"

type GenerateMaskProductRequestDTO struct {
	ProductCode string            `json:"product_code"`
	CreatedBy   uuid.UUID         `json:"created_by"`
	Answers     []MaskAnswerInput `json:"answers"`
}

type MaskAnswerInput struct {
	QuestionID int64 `json:"question_id"`
	OptionID   int64 `json:"option_id"`
	Position   int   `json:"position"`
}
