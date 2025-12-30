package handler

import (
	"encoding/json"
	"net/http"

	"github.com/FelipePn10/panossoerp/internal/application/dto/request"
)

func (h *QuestionHandler) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	var req request.CreateQuestionRequestDTO
	json.NewDecoder(r.Body).Decode(&req)

	err := h.createQuestionUC.Execute(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
