package text

import (
	"bytes"
	"errors"
	"net/http"
)

// Parcel renderer for text data
type Renderer struct{}

// Renders data to an HTTP response writer
func (r Renderer) RenderResponse(w http.ResponseWriter, status int, data interface{}) error {
	s, ok := data.(string)
	if !ok {
		return errors.New("cannot render non-string value")
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(status)

	_, err := w.Write([]byte(s))
	if err != nil {
		return err
	}

	return nil
}

// Creates a new HTTP request and renders data to the body of the request
func (r Renderer) NewRequest(method string, url string, data interface{}) (*http.Request, error) {
	s, ok := data.(string)
	if !ok {
		return nil, errors.New("cannot render non-string value")
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(s)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "text/plain")

	return req, nil
}
