package httpread

import "net/http"

var defaultReader Reader = NewJsonReader()

// Reads from the given HTTP request and writes it to the given data using the default reader
func ReadRequest(r *http.Request, data any) error {
	return defaultReader.ReadRequest(r, data)
}

// Reads from the given HTTP response and writes it to the given data using the default reader
func ReadResponse(r *http.Response, data any) error {
	return defaultReader.ReadResponse(r, data)
}

// Sets the default reader to use for reading requests/responses
func SetDefault(reader Reader) {
	defaultReader = reader
}
