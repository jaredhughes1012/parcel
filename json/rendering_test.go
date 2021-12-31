package json

import (
	"encoding/json"
	"mime"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Renderer_RenderResponse(t *testing.T) {
	type testData struct {
		Field string `json:"field"`
	}

	td1 := testData{Field: "test"}
	var td2 testData

	w := httptest.NewRecorder()
	renderer := Renderer{}

	renderer.RenderResponse(w, http.StatusOK, &td1)
	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("Status code invalid: %d != %d", w.Result().StatusCode, http.StatusOK)
	}

	if ct, _, err := mime.ParseMediaType(w.Header().Get("Content-Type")); err != nil {
		t.Fatal(err)
	} else if ct != "application/json" {
		t.Errorf("Content type invalid: %s != %s", ct, "application/json")
	}

	if err := json.NewDecoder(w.Body).Decode(&td2); err != nil {
		t.Fatal(err)
	} else if td1.Field != td2.Field {
		t.Errorf("Rendering failed: %s != %s", td1.Field, td2.Field)
	}
}

func Test_Renderer_NewRequest(t *testing.T) {
	type testData struct {
		Field string `json:"field"`
	}

	td1 := testData{Field: "test"}
	var td2 testData

	renderer := Renderer{}
	req, err := renderer.NewRequest(http.MethodPost, "/", &td1)
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
