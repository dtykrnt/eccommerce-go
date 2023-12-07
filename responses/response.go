package responses

import "os"

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitepty"`
}

func NewSuccessResponse(message string, data interface{}) *Response {
	prefix := os.Getenv("PROJECT")
	return &Response{
		Success: true,
		Message: prefix + ": " + message,
		Data:    data,
	}
}

func NewErrorResponse(message string) *Response {
	prefix := os.Getenv("PROJECT")
	return &Response{
		Success: false,
		Message: prefix + ": " + message,
	}
}
