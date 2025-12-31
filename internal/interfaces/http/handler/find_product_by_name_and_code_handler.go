package handler

import (
	"encoding/json"
	"net/http"
)

func (h *ProductHandler) FindByNameAndCodeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	name := r.URL.Query().Get("name")
	codeStr := r.URL.Query().Get("code")

	if name == "" || codeStr == "" {
		http.Error(w, "name and code are required", http.StatusBadRequest)
		return
	}

	product, err := h.findProductByNameAndCodeUC.Execute(ctx, name, codeStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}
