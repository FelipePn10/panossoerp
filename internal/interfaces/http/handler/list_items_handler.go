package handler

import (
	"errors"
	"net/http"

	errorsuc "github.com/FelipePn10/panossoerp/internal/application/usecase/errors"
	"github.com/FelipePn10/panossoerp/internal/interfaces/http/handler/security"
)

func (h *ItemHandler) ListItems(w http.ResponseWriter, r *http.Request) {
	items, err := h.listItemsUC.Execute(r.Context())
	if err != nil {
		if errors.Is(err, errorsuc.ErrUnauthorized) {
			security.WriteError(w, http.StatusUnauthorized, "unauthorized", err.Error())
			return
		}
		h.InternalError(w, r, err)
		return
	}
	h.OK(w, items, "items retrieved")
}

func (h *ItemHandler) ListItemsWithMasks(w http.ResponseWriter, r *http.Request) {
	items, err := h.listItemsWithMasksUC.Execute(r.Context())
	if err != nil {
		if errors.Is(err, errorsuc.ErrUnauthorized) {
			security.WriteError(w, http.StatusUnauthorized, "unauthorized", err.Error())
			return
		}
		h.InternalError(w, r, err)
		return
	}
	h.OK(w, items, "items with masks retrieved")
}
