package parcel

import (
	"net/http"

	"github.com/jaredhughes1012/parcel/json"
	"github.com/jaredhughes1012/parcel/text"
)

var (
	textRenderer   = text.Renderer{}
	objectRenderer = json.Renderer{}
)

// Renders data from a source to a destination
type Renderer interface {
	// Renders data to an HTTP response writer
	RenderResponse(w http.ResponseWriter, status int, data interface{}) error

	// Creates a new HTTP request and renders data to the body of the request
	NewRequest(method, destination string, data interface{}) (*http.Request, error)
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
// for string data or the configured object renderer for any other data
func NewRequest(method, u string, data interface{}) (*http.Request, error) {
	_, ok := data.(string)
	if ok {
		return textRenderer.NewRequest(method, u, data)
	} else {
		return objectRenderer.NewRequest(method, u, data)
	}
}
