package json

import (
	"encoding/json"
	"net/http"
)

// Parcel binder for JSON
type Binder struct{}

// Binds data from an HTTP request into a destination
func (binder Binder) BindFromRequest(r *http.Request, target interface{}) error {
	return json.NewDecoder(r.Body).Decode(target)
}

// Binds data from an HTTP response into a destination
func (binder Binder) BindFromResponse(r *http.Response, target interface{}) error {
	return json.NewDecoder(r.Body).Decode(target)
}
