// Handles reading data from HTTP requests
package httpread

import "net/http"

// Reads data from HTTP requests/responses and binds it to data
type Reader interface {
	// Reads from an HTTP request
	ReadRequest(r *http.Request, data any) error

	// Reads from an HTTP response
	ReadResponse(r *http.Response, data any) error
}
