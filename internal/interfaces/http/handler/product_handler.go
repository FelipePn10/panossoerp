package handler

import (
	"encoding/json"
	"net/http"

	"github.com/FelipePn10/panossoerp/internal/application/dto"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateProductDTO
	json.NewDecoder(r.Body).Decode(&req)

	err := h.createProductUC.Execute(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
