package httperror

import (
	"errors"
	"fmt"
	"net/http"
)

// Wrapper around an error with an attached HTTP error code
type HttpError struct {
	Code    int
	Wrapped error
}

// If the given error is an HTTP error, returns the HTTP status code and the wrapped error
// If not, returns the same error and an InternalServerError status (500)
// If the error is nil, returns nil and -1
func Unwrap(err error) (int, error) {
	if err == nil {
		return -1, nil
	}

	if httpErr, ok := err.(*HttpError); !ok {
		return http.StatusInternalServerError, err
	} else {
		return httpErr.Code, httpErr.Wrapped
	}
}

func (err HttpError) Error() string {
	return fmt.Sprintf("http error %d: %s", err.Code, err.Wrapped.Error())
}

// Creates a new error with a wrapped HTTP status code with a formatted message
func Fmt(code int, format string, args ...any) *HttpError {
	return Wrap(code, fmt.Errorf(format, args...))
}

// Creates a new error with a wrapped HTTP status code
func New(code int, message string) *HttpError {
	return Wrap(code, errors.New(message))
}

// Wraps an existing error with an HTTP status code
func Wrap(code int, err error) *HttpError {
	return &HttpError{
		Code:    code,
		Wrapped: err,
	}
}
