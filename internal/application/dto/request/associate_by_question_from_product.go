package request

type AssociateByQuestionProductRequestDTO struct {
	ProductID  int64 `json:"product_id"`
	QuestionID int64 `json:"question_id"`
	Position   int   `json:"position"`
}
