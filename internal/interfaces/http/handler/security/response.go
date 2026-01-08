package security

type SuccessResponse struct {
	Data    any    `json:"data"`
	Message string `json:"message,omitempty"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Details any    `json:"details,omitempty"`
}
