package parcel

import (
	"encoding/json"
	"io"
	"mime"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_RenderResponse_Text(t *testing.T) {
	d1 := "test"

	w := httptest.NewRecorder()
	err := RenderResponse(w, http.StatusCreated, d1)
	if err != nil {
		t.Fatal(err)
	}

	if w.Result().StatusCode != http.StatusCreated {
		t.Errorf("Status code invalid: %d != %d", w.Result().StatusCode, http.StatusCreated)
	}

	if ct, _, err := mime.ParseMediaType(w.Header().Get("Content-Type")); err != nil {
		t.Fatal(err)
	} else if ct != "text/plain" {
		t.Errorf("Content type invalid: %s != %s", ct, "text/plain")
	}

	if b, err := io.ReadAll(w.Body); err != nil {
		t.Fatal(err)
	} else if d2 := string(b); d1 != d2 {
		t.Errorf("Data mismatch: %s != %s", d1, d2)
	}
}

func Test_RenderResponse_Json(t *testing.T) {
	type testData struct {
		Data string `json:"data"`
	}

	d1 := testData{Data: "test"}

	w := httptest.NewRecorder()
	err := RenderResponse(w, http.StatusCreated, &d1)
	if err != nil {
		t.Fatal(err)
	}

	if w.Result().StatusCode != http.StatusCreated {
		t.Errorf("Status code invalid: %d != %d", w.Result().StatusCode, http.StatusCreated)
	}

	if ct, _, err := mime.ParseMediaType(w.Header().Get("Content-Type")); err != nil {
		t.Fatal(err)
	} else if ct != "application/json" {
		t.Errorf("Content type invalid: %s != %s", ct, "application/json")
	}

	var d2 testData
	if err := json.NewDecoder(w.Body).Decode(&d2); err != nil {
		t.Fatal(err)
	} else if d1.Data != d2.Data {
		t.Errorf("Data mismatch: %s != %s", d1.Data, d2.Data)
	}
}

func Test_NewRequest_Object(t *testing.T) {
	type testData struct {
		Field string `json:"field"`
	}

	td1 := testData{Field: "test"}
	var td2 testData

	req, err := NewRequest(http.MethodPost, "/", &td1)
	if err != nil {
		t.Fatal(err)
	}

	if ct, _, err := mime.ParseMediaType(req.Header.Get("Content-Type")); err != nil {
		t.Fatal(err)
	} else if ct != "application/json" {
		t.Errorf("Content type invalid: %s != %s", ct, "application/json")
	}

	if err := json.NewDecoder(req.Body).Decode(&td2); err != nil {
		t.Fatal(err)
	} else if td1.Field != td2.Field {
		t.Errorf("Rendering failed: %s != %s", td1.Field, td2.Field)
	}
}

func Test_NewRequest_Text(t *testing.T) {
	td1 := "test"

	req, err := NewRequest(http.MethodPost, "/", td1)
	if err != nil {
		t.Fatal(err)
	}

	if ct, _, err := mime.ParseMediaType(req.Header.Get("Content-Type")); err != nil {
		t.Fatal(err)
	} else if ct != "text/plain" {
		t.Errorf("Content type invalid: %s != %s", ct, "text/plain")
	}

	b, err := io.ReadAll(req.Body)
	if err != nil {
		t.Fatal(err)
	} else if s := string(b); s != td1 {
		t.Errorf("Rendered value invalid: %s != %s", td1, s)
	}
}

func Test_NewRequest_Nil(t *testing.T) {
	req, err := NewRequest(http.MethodPost, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	if ct := req.Header.Get("Content-Type"); ct != "" {
		t.Errorf("Content type %s should not exist", ct)
	}

	if req.Body != nil {
		t.Error("Request should not have a body")
	}
}
