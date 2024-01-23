package httpwrite

import (
	"context"
	"net/http"
)

// Handles writing data to HTTP requests/responses
type Writer interface {
	// Writes the given data to an HTTP response
	WriteResponse(w http.ResponseWriter, status int, data any) error

	// Writes the given error to an HTTP response
	WriteError(w http.ResponseWriter, err error)

	// Creates a new HTTP request and encodes the given data as the body of the request
	NewRequest(ctx context.Context, method, url string, data any) (*http.Request, error)
}
