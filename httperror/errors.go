package httperror

import (
	"fmt"
	"net/http"
)

// Error containing an internal error and corresponding http response data
type HttpError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	wrapped error  `json:"-"`
}

func (err HttpError) Error() string {
	return fmt.Sprintf("http error %d: %s", err.Status, err.Message)
}

// Creates a new error with a format string
func New(status int, format string, args ...interface{}) error {
	wrapped := fmt.Errorf(format, args...)
	return &HttpError{
		Status:  status,
		wrapped: wrapped,
		Message: wrapped.Error(),
	}
}

// Wraps an existing error with an http status code
func Wrap(status int, wrapped error) error {
	return &HttpError{
		Status:  status,
		wrapped: wrapped,
		Message: wrapped.Error(),
	}
}

// Checks if the error is an http error. If it is, returns it cast as an HttpError
// If not, provides a new wrapped http error
func Convert(err error) *HttpError {
	he, ok := err.(*HttpError)
	if !ok {
		he = &HttpError{
			Status:  http.StatusInternalServerError,
			wrapped: err,
			Message: err.Error(),
		}
	}

	return he
}
