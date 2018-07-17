package model

// ErrorResponse represents the error response body
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
