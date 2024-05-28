package exception

import "encoding/json"

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (error *ErrorResponse) Error() string {
	log := &ErrorResponse{
		Code:    error.Code,
		Message: error.Message,
	}

	jsonBytes, _ := json.Marshal(log)
	return string(jsonBytes)
}

func NewErrorResponse(code int, message string) *ErrorResponse {
	return &ErrorResponse{
		Code:    code,
		Message: message,
	}
}
