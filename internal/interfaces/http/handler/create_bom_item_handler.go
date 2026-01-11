package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/FelipePn10/panossoerp/internal/application/dto/request"
	errorsuc "github.com/FelipePn10/panossoerp/internal/application/usecase/errors"
)

func (h *BomItemHandler) Create(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req request.CreateBomItemsRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.BadRequest(w, "invalid request body")
		return
	}

	bomitem, err := h.createBomItemUC.Execute(r.Context(), req)
	if err != nil {
		switch {
		case errors.Is(err, errorsuc.ErrCreateBomItem):
			h.BadRequest(w, "falied create bom item")
		case errors.Is(err, errorsuc.ErrCreateBomItemNotFound):
			h.NotFound(w, "try again later")
		default:
			h.InternalError(w, err)
			return
		}
		return
	}
	h.OK(w, bomitem, "Create Bom Item Success!")
}
