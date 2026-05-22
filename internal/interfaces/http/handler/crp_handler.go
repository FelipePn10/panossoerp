package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/FelipePn10/panossoerp/internal/application/dto/request"
	"github.com/go-chi/chi/v5"
)

func (h *CRPHandler) CalculateCRP(w http.ResponseWriter, r *http.Request) {
	var dto request.CalculateCRPDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		jsonError(w, http.StatusBadRequest, "invalid payload: "+err.Error())
		return
	}
	result, err := h.uc.CalculateCRP(r.Context(), dto)
	if err != nil {
		jsonError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	jsonResponse(w, http.StatusOK, result)
}

func (h *CRPHandler) ListByPlan(w http.ResponseWriter, r *http.Request) {
	planCode, err := strconv.ParseInt(chi.URLParam(r, "planCode"), 10, 64)
	if err != nil {
		jsonError(w, http.StatusBadRequest, "invalid planCode")
		return
	}
	result, err := h.uc.ListByPlan(r.Context(), planCode)
	if err != nil {
		jsonError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	jsonResponse(w, http.StatusOK, result)
}

func (h *CRPHandler) ListOverloadedByPlan(w http.ResponseWriter, r *http.Request) {
	planCode, err := strconv.ParseInt(chi.URLParam(r, "planCode"), 10, 64)
	if err != nil {
		jsonError(w, http.StatusBadRequest, "invalid planCode")
		return
	}
	result, err := h.uc.ListOverloadedByPlan(r.Context(), planCode)
	if err != nil {
		jsonError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	jsonResponse(w, http.StatusOK, result)
}
