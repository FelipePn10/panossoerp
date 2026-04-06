package request

type AssociateProductQuestionsRequest struct {
	ItemID    int64 `json:"item_id"`
	Questions []struct {
		QuestionID int64 `json:"question_id"`
		Position   int   `json:"position"`
	} `json:"questions"`
}
