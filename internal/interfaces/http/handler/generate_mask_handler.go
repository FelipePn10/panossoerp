package handler

import (
	"encoding/json"
	"net/http"

	applicationreq "github.com/FelipePn10/panossoerp/internal/application/dto/request"
	internalreq "github.com/FelipePn10/panossoerp/internal/infrastructure/http/request"
)

func (h *GenerateMaskHandler) GenerateMask(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req internalreq.GenerateMaskProduct
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.ProductCode == "" || len(req.Answers) == 0 {
		http.Error(w, "productCode and answers are required", http.StatusBadRequest)
		return
	}

	answers := make([]applicationreq.MaskAnswerInput, 0, len(req.Answers))
	for _, a := range req.Answers {
		answers = append(answers, applicationreq.MaskAnswerInput{
			QuestionID: a.QuestionID,
			OptionID:   a.OptionID,
			Position:   a.Position,
		})
	}
	d := applicationreq.GenerateMaskProductRequestDTO{
		ProductCode: req.ProductCode,
		Answers:     answers,
	}

	if err := h.generateMask.Execute(r.Context(), d); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
