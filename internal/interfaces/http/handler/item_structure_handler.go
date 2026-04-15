package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/FelipePn10/panossoerp/internal/application/dto/request"
	"github.com/go-chi/chi/v5"
)

func (h *ItemStructureHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto request.CreateStructureComponentDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		jsonError(w, http.StatusBadRequest, "payload inválid usecase.NewCreateStructureComponentUseCase(itemRepoStructure, authService)\n\tupdateStructureUc := usecase.NewUpdateStructureCompoo: "+err.Error())
		return
	}

	result, err := h.createUC.Execute(r.Context(), dto)
	if err != nil {
		jsonError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	jsonResponse(w, http.StatusCreated, result)
}

func (h *ItemStructureHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r, "id")
	if err != nil {
		jsonError(w, http.StatusBadRequest, err.Error())
		return
	}

	var dto request.UpdateStructureComponentDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		jsonError(w, http.StatusBadRequest, "invalid payload: "+err.Error())
		return
	}
	dto.ID = id

	result, err := h.updateUC.Execute(r.Context(), dto)
	if err != nil {
		jsonError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, result)
}

func (h *ItemStructureHandler) Delete(w http.ResponseWriter, r *http.Request) {
	//id
	_, err := parseID(r, "id")
	if err != nil {
		jsonError(w, http.StatusBadRequest, err.Error())
		return
	}

	//if err := h.deleteUC.Execute(r.Context(), id); err != nil {
	//	jsonError(w, http.StatusUnprocessableEntity, err.Error())
	//	return
	//}

	w.WriteHeader(http.StatusNoContent)
}

// Retorna a árvore BOM genérica de um item (sem resolução de máscara).
func (h *ItemStructureHandler) GetTree(w http.ResponseWriter, r *http.Request) {
	rootItemID, err := parseID(r, "rootItemId")
	if err != nil {
		jsonError(w, http.StatusBadRequest, err.Error())
		return
	}

	dto := request.GetStructureTreeDTO{RootItemID: rootItemID}
	result, err := h.treeUC.Execute(r.Context(), dto)
	if err != nil {
		jsonError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, result)
}

// Retorna todos os filhos diretos de um item.
func (h *ItemStructureHandler) GetAllDirectChildren(w http.ResponseWriter, r *http.Request) {
	parentItemID, err := parseID(r, "parentItemId")
	if err != nil {
		jsonError(w, http.StatusBadRequest, err.Error())
		return
	}

	dto := request.GetAllDirectChildrenDTO{
		ParentItemID: parentItemID,
	}

	result, err := h.getAllStructure.Execute(r.Context(), dto)
	if err != nil {
		jsonError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, result)
}

// Resolve a árvore BOM completa para uma configuração (máscara) específica.
// Query param: mask (ex.: "100#100#50" — encode o # como %23 na URL)
func (h *ItemStructureHandler) ResolveForMask(w http.ResponseWriter, r *http.Request) {
	rootItemID, err := parseID(r, "rootItemId")
	if err != nil {
		jsonError(w, http.StatusBadRequest, err.Error())
		return
	}

	maskValue := r.URL.Query().Get("mask")
	if maskValue == "" {
		jsonError(w, http.StatusBadRequest, "query param 'mask' is required (ex.: ?mask=100%23100%2350)")
		return
	}

	dto := request.ResolveStructureForMaskDTO{
		RootItemID:    rootItemID,
		RootMaskValue: maskValue,
	}

	result, err := h.resolveUC.Execute(r.Context(), dto)
	if err != nil {
		jsonError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, result)
}

func parseID(r *http.Request, param string) (int64, error) {
	raw := chi.URLParam(r, param)
	id, err := strconv.ParseInt(raw, 10, 64)
	if err != nil || id <= 0 {
		return 0, fmt.Errorf("the parameter '%s' must be a positive integer", param)
	}
	return id, nil
}

func jsonResponse(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

func jsonError(w http.ResponseWriter, status int, msg string) {
	jsonResponse(w, status, map[string]string{"error": msg})
}
