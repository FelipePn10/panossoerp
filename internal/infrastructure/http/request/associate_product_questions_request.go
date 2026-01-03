package request

type AssociateProductQuestionsRequest struct {
	ProductID int64 `json:"product_id"`
	Questions []struct {
		QuestionID int64 `json:"question_id"`
		Position   int   `json:"position"`
	} `json:"questions"`
}
