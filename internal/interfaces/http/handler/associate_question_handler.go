package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	applicationreq "github.com/FelipePn10/panossoerp/internal/application/dto/request"
	internalreq "github.com/FelipePn10/panossoerp/internal/infrastructure/http/request"
	"github.com/FelipePn10/panossoerp/internal/interfaces/http/handler/security"
	"github.com/go-chi/chi/v5"
)

func (h *AssociateByQuestionItemHandler) AssociateQuestions(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req internalreq.AssociateProductQuestionsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.ItemCode <= 0 {
		http.Error(w, "item_id is required and must be greater than zero", http.StatusBadRequest)
		return
	}
	if len(req.Questions) == 0 {
		http.Error(w, "questions cannot be empty", http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	// one product N questions
	for _, q := range req.Questions {
		dto := applicationreq.AssociateByQuestionItemRequestDTO{
			ItemCode:   req.ItemCode,
			QuestionID: q.QuestionID,
			Position:   q.Position,
		}
		if err := h.associateByQuestionProductUC.Execute(ctx, dto); err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *AssociateByQuestionItemHandler) GetQuestionsByItem(w http.ResponseWriter, r *http.Request) {
	itemCode, err := strconv.ParseInt(chi.URLParam(r, "itemCode"), 10, 64)
	if err != nil || itemCode <= 0 {
		security.WriteError(w, http.StatusBadRequest, "bad_request", "invalid item_code")
		return
	}

	questions, err := h.getQuestionsByItemUC.Execute(r.Context(), itemCode)
	if err != nil {
		h.InternalError(w, r, err)
		return
	}

	h.OK(w, questions, "questions retrieved")
}

func (h *AssociateByQuestionItemHandler) ListAll(w http.ResponseWriter, r *http.Request) {
	rows, err := h.listAllItemQuestionsUC.Execute(r.Context())
	if err != nil {
		h.InternalError(w, r, err)
		return
	}

	h.OK(w, rows, "all item questions retrieved")
}
