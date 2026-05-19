package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/FelipePn10/panossoerp/internal/application/dto/request"
	errorsuc "github.com/FelipePn10/panossoerp/internal/application/usecase/errors"
	"github.com/FelipePn10/panossoerp/internal/interfaces/http/handler/security"
)

func (h *QuestionOptionHandler) CreateQuestionOptionHandler(w http.ResponseWriter, r *http.Request) {
	var req request.CreateQuestionOptionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		security.WriteError(w, http.StatusBadRequest, "bad_request", "invalid request body")
		return
	}

	questionOption, err := h.createQuestionOptionUC.Execute(r.Context(), req)
	if err != nil {
		switch {
		case errors.Is(err, errorsuc.ErrUnauthorized):
			security.WriteError(w, http.StatusUnauthorized, "unauthorized", err.Error())
		case errors.Is(err, errorsuc.ErrQuestionOptionAlreadyExists):
			security.WriteError(w, http.StatusConflict, "conflict", err.Error())
		default:
			h.InternalError(w, r, err)
		}
		return
	}
	h.OK(w, questionOption, "Created question option success")
}
