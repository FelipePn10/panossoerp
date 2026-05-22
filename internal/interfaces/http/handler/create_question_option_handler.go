package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/FelipePn10/panossoerp/internal/application/dto/request"
	errorsuc "github.com/FelipePn10/panossoerp/internal/application/usecase/errors"
	"github.com/FelipePn10/panossoerp/internal/interfaces/http/handler/security"
	"github.com/go-chi/chi/v5"
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

func (h *QuestionOptionHandler) ListByQuestion(w http.ResponseWriter, r *http.Request) {
	questionID, err := strconv.ParseInt(chi.URLParam(r, "questionID"), 10, 64)
	if err != nil || questionID <= 0 {
		security.WriteError(w, http.StatusBadRequest, "bad_request", "invalid question_id")
		return
	}

	options, err := h.listOptionsByQuestionUC.Execute(r.Context(), questionID)
	if err != nil {
		if errors.Is(err, errorsuc.ErrUnauthorized) {
			security.WriteError(w, http.StatusUnauthorized, "unauthorized", err.Error())
			return
		}
		h.InternalError(w, r, err)
		return
	}

	h.OK(w, options, "options retrieved")
}
