package httpwrite

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jaredhughes1012/parcel/httperror"
	"github.com/stretchr/testify/assert"
)

func Test_JsonWriter_WriteError(t *testing.T) {
	cases := []struct {
		name           string
		err            error
		expectedStatus int
	}{
		{
			name:           "standard error",
			err:            errors.New("test error"),
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "http error",
			err:            httperror.New(http.StatusBadRequest, "test error"),
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			writer := NewJsonWriter()

			writer.WriteError(w, c.err)
			assert.Equal(t, c.expectedStatus, w.Code)

			var result map[string]any
			assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
			assert.NoError(t, json.NewDecoder(w.Body).Decode(&result))

			// Test as default
			w = httptest.NewRecorder()
			SetDefaultWriter(writer)

			Error(w, c.err)
			assert.Equal(t, c.expectedStatus, w.Code)
			assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
			assert.NoError(t, json.NewDecoder(w.Body).Decode(&result))
		})
	}
}

func Test_JsonWriter_WriteRequest(t *testing.T) {
	data := map[string]any{
		"test": "value",
	}

	w := httptest.NewRecorder()
	writer := NewJsonWriter()

	err := writer.WriteResponse(w, http.StatusBadRequest, data)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var result map[string]any
	err = json.NewDecoder(w.Body).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, data["test"], result["test"])

	// Test as default

	w = httptest.NewRecorder()
	SetDefaultWriter(writer)

	err = Response(w, http.StatusNotFound, data)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	err = json.NewDecoder(w.Body).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, data["test"], result["test"])
}

func Test_JsonWriter_NewRequest(t *testing.T) {
	data := map[string]any{
		"test": "value",
	}

	writer := NewJsonWriter()
	r, err := writer.NewRequest(context.Background(), http.MethodPut, "/", data)
	assert.NoError(t, err)
	assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

	var result map[string]any
	err = json.NewDecoder(r.Body).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, data["test"], result["test"])
}
