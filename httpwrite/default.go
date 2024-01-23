package httpwrite

import "net/http"

var defaultWriter Writer = NewJsonWriter()

// SetDefaultWriter sets the default writer to use for writing HTTP responses.
func SetDefaultWriter(writer Writer) {
	defaultWriter = writer
}

// Response writes the given data to an HTTP response using the default writer.
func Response(w http.ResponseWriter, status int, data any) error {
	return defaultWriter.WriteResponse(w, status, data)
}

// Error writes the given error to an HTTP response using the default writer.
func Error(w http.ResponseWriter, err error) {
	defaultWriter.WriteError(w, err)
}
