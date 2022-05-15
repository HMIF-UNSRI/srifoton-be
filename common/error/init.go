package error

import (
	"fmt"
	"net/http"
)

type ClientError struct {
	StatusCode int
	Message    string
}

var _ error = &ClientError{}

func (e ClientError) Error() string {
	return fmt.Sprintf("%d\t%s", e.StatusCode, e.Message)
}

func NewInvariantError(msg string) *ClientError {
	return &ClientError{
		StatusCode: http.StatusBadRequest,
		Message:    msg,
	}
}

func NewNotFoundError(msg string) *ClientError {
	return &ClientError{
		StatusCode: http.StatusNotFound,
		Message:    msg,
	}
}

func NewForbiddenError(msg string) *ClientError {
	return &ClientError{
		StatusCode: http.StatusForbidden,
		Message:    msg,
	}
}

func NewUnauthorizedError(msg string) *ClientError {
	return &ClientError{
		StatusCode: http.StatusUnauthorized,
		Message:    msg,
	}
}
