package request

import "github.com/google/uuid"

type GenerateMaskItem struct {
	ItemCode  string    `json:"item_code"`
	ItemID    int64     `json:"item_id"`
	CreatedBy uuid.UUID `json:"created_by"`
	Answers   []struct {
		QuestionID int64 `json:"question_id"`
		OptionID   int64 `json:"option_id"`
		Position   int   `json:"position"`
	} `json:"answers"`
}
