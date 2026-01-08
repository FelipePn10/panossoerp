package security

import (
	"encoding/json"
	"net/http"
)

func WriteSuccess(
	w http.ResponseWriter,
	status int,
	data any,
	message string,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := SuccessResponse{
		Data:    data,
		Message: message,
	}
	_ = json.NewEncoder(w).Encode(resp)
}

func WriteError(
	w http.ResponseWriter,
	status int,
	errKey,
	message string,
	details ...any,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := ErrorResponse{
		Error:   errKey,
		Message: message,
	}

	if len(details) > 0 {
		resp.Details = details[0]
	}

	_ = json.NewEncoder(w).Encode(resp)
}
