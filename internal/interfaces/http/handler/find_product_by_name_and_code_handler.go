package handler

import (
	"errors"
	"net/http"
	"strings"

	errorsuc "github.com/FelipePn10/panossoerp/internal/application/usecase/errors"
)

func (h *ProductHandler) FindByNameAndCodeHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	name := strings.TrimSpace(r.URL.Query().Get("name"))
	codeStr := strings.TrimSpace(r.URL.Query().Get("codeStr"))

	if name == "" || codeStr == "" {
		h.BadRequest(w, "name and code are required")
		return
	}

	product, err := h.findProductByNameAndCodeUC.Execute(r.Context(), name, codeStr)
	if err != nil {
		switch {
		case errors.Is(err, errorsuc.ErrInvalidProductNameAndCodeNotFound):
			h.BadRequest(w, "The 'name' and 'code' query parameter is required")
			return
		case errors.Is(err, errorsuc.ErrProductNotFound):
			h.NotFound(w, "No product found with the given name and code")
			return
		default:
			h.InternalError(w, err)
			return
		}
	}

	h.OK(w, product, "Product Found")
}
