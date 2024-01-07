package httpwrite

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/jaredhughes1012/parcel/httperror"
)

type JsonWriter struct{}

// WriteError implements Writer.
func (writer JsonWriter) WriteError(w http.ResponseWriter, err error) {
	status, _ := httperror.Unwrap(err)

	w.WriteHeader(status)
	_ = writer.WriteResponse(w, map[string]any{
		"error": http.StatusText(status),
	})
}

// NewRequest implements Writer.
func (JsonWriter) NewRequest(ctx context.Context, method, url string, data any) (*http.Request, error) {
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

// WriteResponse implements Writer.
func (JsonWriter) WriteResponse(w http.ResponseWriter, data any) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(data)
}

var _ Writer = (*JsonWriter)(nil)

func NewJsonWriter() *JsonWriter {
	return &JsonWriter{}
}
