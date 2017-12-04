package errors

import (
	"encoding/json"
	"net/http"
)

type HTTPError struct {
	Code    int
	Message string
}

func NewHTTPError(code int, message string) *HTTPError {
	return &HTTPError{
		Code:    code,
		Message: message,
	}
}

func (e *HTTPError) WriteTo(w http.ResponseWriter) {
	w.WriteHeader(e.Code)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error_code":    e.Code,
		"error_message": e.Message,
	})
}
