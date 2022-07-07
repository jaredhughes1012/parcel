package parcel

import (
	"net/http"

	"github.com/jaredhughes1012/parcel/json"
	"github.com/jaredhughes1012/parcel/text"
)

var (
	textRenderer   = Renderer(text.Renderer{})
	objectRenderer = Renderer(json.Renderer{})
	errorRenderer  = Renderer(json.Renderer{})
)

// Sets the default renderer for any plain text responses
func SetTextRenderer(r Renderer) {
	textRenderer = r
}

// Sets the default renderer for any error responses
func SetErrorRenderer(r Renderer) {
	errorRenderer = r
}

// Sets the default renderer for any object responses
func SetObjectRenderer(r Renderer) {
	objectRenderer = r
}

// Renders data from a source to a destination
type Renderer interface {
	// Renders data to an HTTP response writer
	RenderResponse(w http.ResponseWriter, status int, data interface{}) error

	// Creates a new HTTP request and renders data to the body of the request
	NewRequest(method, destination string, data interface{}) (*http.Request, error)

	// Renders an error to an HTTP response
	RenderErrorResponse(w http.ResponseWriter, err error)
}

// Renders data from the given data into an HTTP response. Will use the configured text renderer
// for string data or the configured object renderer for any other data
func RenderResponse(w http.ResponseWriter, status int, data interface{}) error {
	_, ok := data.(string)
	if ok {
		return textRenderer.RenderResponse(w, status, data)
	} else {
		return objectRenderer.RenderResponse(w, status, data)
	}
}

// Creates a new HTTP request and renders the given data into it. Will use the configured text renderer
// for string data or the configured object renderer for any other data. If data is nil, this will behave
// the same as using http.NewRequest with a nil body
func NewRequest(method, u string, data interface{}) (*http.Request, error) {
	if data == nil {
		return http.NewRequest(method, u, nil)
	}

	_, ok := data.(string)
	if ok {
		return textRenderer.NewRequest(method, u, data)
	} else {
		return objectRenderer.NewRequest(method, u, data)
	}
}

// Renders an error into an HTTP response. Uses JSON formatting by default but can be configured to use any
// renderer. Will use the status of any HTTP error or InternalServerError(500) if the error is not an http
// error
func RenderErrorResponse(w http.ResponseWriter, err error) {
	errorRenderer.RenderErrorResponse(w, err)
}
