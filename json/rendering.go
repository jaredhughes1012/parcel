package json

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/jaredhughes1012/parcel/httperror"
)

// Parcel renderer for json data
type Renderer struct{}

// Renders data to an HTTP response writer
func (r Renderer) RenderResponse(w http.ResponseWriter, status int, data interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, err = w.Write(b)
	if err != nil {
		return err
	}

	return nil
}

// Creates a new HTTP request and renders data to the body of the request
func (r Renderer) NewRequest(method string, url string, data interface{}) (*http.Request, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (r Renderer) RenderErrorResponse(w http.ResponseWriter, err error) {
	he := httperror.Convert(err)
	_ = r.RenderResponse(w, he.Status, he)
}
