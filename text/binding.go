package text

import (
	"errors"
	"io"
	"net/http"
)

// Parcel binder for plain text
type Binder struct{}

// Binds data from an HTTP request into a destination
func (binder Binder) BindFromRequest(r *http.Request, target interface{}) error {
	tp, ok := target.(*string)
	if !ok {
		return errors.New("invalid target, must be a string pointer")
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	*tp = string(b)
	return nil
}

// Binds data from an HTTP response into a destination
func (binder Binder) BindFromResponse(r *http.Response, target interface{}) error {
	tp, ok := target.(*string)
	if !ok {
		return errors.New("invalid target, must be a string pointer")
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	*tp = string(b)
	return nil
}
