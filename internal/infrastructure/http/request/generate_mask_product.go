package request

import "github.com/google/uuid"

type GenerateMaskItem struct {
	ItemCode  int64     `json:"item_code"`
	CreatedBy uuid.UUID `json:"created_by"`
	Answers   []struct {
		QuestionID int64 `json:"question_id"`
		OptionID   int64 `json:"option_id"`
		Position   int   `json:"position"`
	} `json:"answers"`
}
