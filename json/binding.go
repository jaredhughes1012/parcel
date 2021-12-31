package json

import (
	"encoding/json"
	"net/http"

	"github.com/jaredhughes1012/parcel/httperror"
)

// Parcel binder for JSON
type Binder struct{}

// Binds data from an HTTP request into a destination
func (binder Binder) BindFromRequest(r *http.Request, target interface{}) error {
	err := json.NewDecoder(r.Body).Decode(target)
	if err != nil {
		return httperror.Wrap(http.StatusBadRequest, err)
	}

	return nil
}

// Binds data from an HTTP response into a destination
func (binder Binder) BindFromResponse(r *http.Response, target interface{}) error {
	err := json.NewDecoder(r.Body).Decode(target)
	if err != nil {
		return httperror.Wrap(http.StatusBadRequest, err)
	}

	return nil
}
