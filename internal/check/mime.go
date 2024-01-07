package check

import (
	"mime"
	"net/http"

	"github.com/jaredhughes1012/parcel/httperror"
)

// Checks if the given content-type header matches the expected content type
func Mime(header, expected string) error {
	if header == "" {
		return httperror.Fmt(http.StatusBadRequest, "no content type provided")
	}

	ct, _, err := mime.ParseMediaType(header)
	if err != nil {
		return err
	} else if ct != expected {
		return httperror.Fmt(http.StatusUnsupportedMediaType, "invalid content type %s", ct)
	}

	return nil
}
