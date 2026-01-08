package security

import (
	"log/slog"
	"net/http"
)

type BaseHandler struct{}

func (h *BaseHandler) OK(w http.ResponseWriter, data any, msg ...string) {
	message := "success"
	if len(msg) > 0 {
		message = msg[0]
	}
	WriteSuccess(w, http.StatusOK, data, message)
}

func (h *BaseHandler) Created(w http.ResponseWriter, data any, msg ...string) {
	message := "created"
	if len(msg) > 0 {
		message = msg[0]
	}
	WriteSuccess(w, http.StatusCreated, data, message)
}

func (h *BaseHandler) BadRequest(w http.ResponseWriter, message string, details ...any) {
	WriteError(w, http.StatusBadRequest, "bad_request", message, details...)
}

func (h *BaseHandler) NotFound(w http.ResponseWriter, message ...string) {
	msg := "resource not found"
	if len(message) > 0 {
		msg = message[0]
	}
	WriteError(w, http.StatusNotFound, "not_found", msg)
}

func (h *BaseHandler) InternalError(w http.ResponseWriter, err error) {
	slog.Error("internal error", "err", err, "path", http.StatusInternalServerError)
	WriteError(w, http.StatusInternalServerError, "internal_error", "Something went wrong")
}
