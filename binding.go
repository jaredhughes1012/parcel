package parcel

import (
	"fmt"
	"mime"
	"net/http"

	"github.com/jaredhughes1012/parcel/json"
	"github.com/jaredhughes1012/parcel/text"
)

var (
	objectBinders = map[string]Binder{
		"application/json": json.Binder{},
	}
	textBinder = text.Binder{}
)

// Handles binding data from a source to a destination
type Binder interface {
	// Binds data from an HTTP request into a destination
	BindFromRequest(r *http.Request, target interface{}) error

	// Binds data from an HTTP response into a destination
	BindFromResponse(r *http.Response, target interface{}) error
}

// Binds data from the request to the given target. Will use a text binder if the target is a string
// pointer, or an object binder otherwise
func BindFromRequest(r *http.Request, target interface{}) error {
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return err
	}

	_, ok := target.(*string)
	if ok {
		if ct == "text/plain" {
			return textBinder.BindFromRequest(r, target)
		} else {
			return fmt.Errorf("cannot bind content type %s to string", ct)
		}
	}

	binder := objectBinders[ct]
	if binder == nil {
		return fmt.Errorf("unsupported content type %s", ct)
	}

	return binder.BindFromRequest(r, target)
}

// Binds data from the request to the given target. Will use a text binder if the target is a string
// pointer, or an object binder otherwise
func BindFromResponse(r *http.Response, target interface{}) error {
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return err
	}

	_, ok := target.(*string)
	if ok {
		if ct == "text/plain" {
			return textBinder.BindFromResponse(r, target)
		} else {
			return fmt.Errorf("cannot bind content type %s to string", ct)
		}
	}

	binder := objectBinders[ct]
	if binder == nil {
		return fmt.Errorf("unsupported content type %s", ct)
	}

	return binder.BindFromResponse(r, target)
}
