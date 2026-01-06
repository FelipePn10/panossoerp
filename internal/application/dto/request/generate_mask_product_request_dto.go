package request

type GenerateMaskProductRequestDTO struct {
	ProductCode string            `json:"product_code"`
	Answers     []MaskAnswerInput `json:"answers"`
}

type MaskAnswerInput struct {
	QuestionID int64 `json:"question_id"`
	OptionID   int64 `json:"option_id"`
	Position   int   `json:"position"`
}
