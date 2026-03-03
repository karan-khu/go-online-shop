package exception

import (
	"fmt"
	"net/http"
)

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

type HTTPError struct {
	StatusCode int
	Body       Error
}

func (e *HTTPError) Error() string {
	return e.Body.Error()
}

func BadRequest(message string) *HTTPError {
	return &HTTPError{
		StatusCode: http.StatusBadRequest,
		Body:       Error{Code: "BAD_REQUEST", Message: message},
	}
}

func NotFound(message string) *HTTPError {
	return &HTTPError{
		StatusCode: http.StatusNotFound,
		Body:       Error{Code: "NOT_FOUND", Message: message},
	}
}

func InternalServerError(message string) *HTTPError {
	return &HTTPError{
		StatusCode: http.StatusInternalServerError,
		Body:       Error{Code: "INTERNAL_ERROR", Message: message},
	}
}
