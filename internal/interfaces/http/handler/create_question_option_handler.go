package handler

import (
	"encoding/json"
	"net/http"

	"github.com/FelipePn10/panossoerp/internal/application/dto/request"
)

func (h *QuestionOptionHandler) CreateQuestionOptionHandler(w http.ResponseWriter, r *http.Request) {
	var req request.CreateQuestionOptionRequest
	json.NewDecoder(r.Body).Decode(&req)

	err := h.createQuestionOptionUC.Execute(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
