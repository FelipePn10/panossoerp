package request

import "github.com/google/uuid"

type GenerateMaskItemRequestDTO struct {
	ItemCode  string            `json:"item_code"`
	Answers   []MaskAnswerInput `json:"answers"`
	CreatedBy uuid.UUID         `json:"created_by"`
}

type MaskAnswerInput struct {
	QuestionID int64 `json:"question_id"`
	OptionID   int64 `json:"option_id"`
	Position   int   `json:"position"`
}
