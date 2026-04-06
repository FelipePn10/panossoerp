package request

type AssociateByQuestionItemRequestDTO struct {
	ItemID     int64 `json:"item_id"`
	QuestionID int64 `json:"question_id"`
	Position   int   `json:"position"`
}
