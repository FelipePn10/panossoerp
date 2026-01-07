package handler

import (
	"encoding/json"
	"net/http"
)

func (h *QuestionHandler) FindQuestionByName(
	w http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()
	name := r.URL.Query().Get("name")

	if name == "" {
		http.Error(w, "name required", http.StatusBadRequest)
		return
	}

	question, err := h.findQuestionByNameUC.Execute(ctx, name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(question)
}
