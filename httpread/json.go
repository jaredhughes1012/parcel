package httpread

import (
	"encoding/json"
	"net/http"

	"github.com/jaredhughes1012/parcel/httperror"
	"github.com/jaredhughes1012/parcel/internal/check"
)

type JsonReader struct{}

// ReadResponse implements Reader.
func (*JsonReader) ReadResponse(r *http.Response, data any) error {
	if err := check.Mime(r.Header.Get("Content-Type"), "application/json"); err != nil {
		return err
	} else if err = json.NewDecoder(r.Body).Decode(data); err != nil {
		return httperror.Wrap(http.StatusBadRequest, err)
	}

	return nil
}

// ReadRequest implements Reader.
func (JsonReader) ReadRequest(r *http.Request, data any) error {
	if err := check.Mime(r.Header.Get("Content-Type"), "application/json"); err != nil {
		return err
	} else if err = json.NewDecoder(r.Body).Decode(data); err != nil {
		return httperror.Wrap(http.StatusBadRequest, err)
	}

	return nil
}

var _ Reader = (*JsonReader)(nil)

func NewJsonReader() *JsonReader {
	return &JsonReader{}
}
