package httpread

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jaredhughes1012/parcel/httperror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_JsonReader_ReadRequest_Success(t *testing.T) {
	data := map[string]any{
		"test": "value",
	}

	b, err := json.Marshal(data)
	require.NoError(t, err)

	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(b))
	r.Header.Set("Content-Type", "application/json")

	var result map[string]any
	reader := NewJsonReader()
	err = reader.ReadRequest(r, &result)

	assert.NoError(t, err)
	assert.Equal(t, data["test"], result["test"])

	// Test as default reader
	r = httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(b))
	r.Header.Set("Content-Type", "application/json")

	SetDefault(reader)
	err = ReadRequest(r, &result)

	assert.NoError(t, err)
	assert.Equal(t, data["test"], result["test"])
}

func Test_JsonReader_ReadRequest_NoContentType(t *testing.T) {
	data := map[string]any{
		"test": "value",
	}

	b, err := json.Marshal(data)
	require.NoError(t, err)

	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(b))

	var result map[string]any
	err = NewJsonReader().ReadRequest(r, &result)

	status, err := httperror.Unwrap(err)
	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, status)
}

func Test_JsonReader_ReadRequest_BadContent(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte("text content")))
	r.Header.Set("Content-Type", "text/plain")

	var result map[string]any
	err := NewJsonReader().ReadRequest(r, &result)

	status, err := httperror.Unwrap(err)
	assert.Error(t, err)
	assert.Equal(t, http.StatusUnsupportedMediaType, status)
}

func Test_JsonReader_ReadRequest_BadJson(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte("}{")))
	r.Header.Set("Content-Type", "application/json")

	var result map[string]any
	err := NewJsonReader().ReadRequest(r, &result)

	status, err := httperror.Unwrap(err)
	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, status)
}

func Test_JsonReader_ReadResponse_Success(t *testing.T) {
	data := map[string]any{
		"test": "value",
	}

	b, err := json.Marshal(data)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(b)
	require.NoError(t, err)

	var result map[string]any
	reader := NewJsonReader()
	err = reader.ReadResponse(w.Result(), &result)

	assert.NoError(t, err)
	assert.Equal(t, data["test"], result["test"])

	// Test as default reader
	w = httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(b)
	require.NoError(t, err)

	SetDefault(reader)
	err = ReadResponse(w.Result(), &result)

	assert.NoError(t, err)
	assert.Equal(t, data["test"], result["test"])
}

func Test_JsonReader_ReadResponse_NoContentType(t *testing.T) {
	data := map[string]any{
		"test": "value",
	}

	b, err := json.Marshal(data)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	_, err = w.Write(b)
	require.NoError(t, err)
	r := w.Result()
	r.Header.Set("Content-Type", "")

	var result map[string]any
	reader := NewJsonReader()
	err = reader.ReadResponse(w.Result(), &result)

	status, err := httperror.Unwrap(err)
	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, status)
}

func Test_JsonReader_ReadResponse_BadContent(t *testing.T) {
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "text/plain")
	_, err := w.Write([]byte("text content"))
	require.NoError(t, err)

	var result map[string]any
	err = NewJsonReader().ReadResponse(w.Result(), &result)

	status, err := httperror.Unwrap(err)
	assert.Error(t, err)
	assert.Equal(t, http.StatusUnsupportedMediaType, status)
}

func Test_JsonReader_ReadResponse_BadJson(t *testing.T) {
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write([]byte("}{"))
	require.NoError(t, err)

	var result map[string]any
	err = NewJsonReader().ReadResponse(w.Result(), &result)

	status, err := httperror.Unwrap(err)
	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, status)
}
