package text

import (
	"io"
	"mime"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Renderer_RenderResponse(t *testing.T) {
	td1 := "test"

	w := httptest.NewRecorder()
	renderer := Renderer{}

	renderer.RenderResponse(w, http.StatusOK, td1)
	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("Status code invalid: %d != %d", w.Result().StatusCode, http.StatusOK)
	}

	if ct, _, err := mime.ParseMediaType(w.Header().Get("Content-Type")); err != nil {
		t.Fatal(err)
	} else if ct != "text/plain" {
		t.Errorf("Content type invalid: %s != %s", ct, "text/plain")
	}

	b, err := io.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	} else if s := string(b); s != td1 {
		t.Errorf("Rendered value invalid: %s != %s", td1, s)
	}
}

func Test_Renderer_NewRequest(t *testing.T) {
	td1 := "test"

	renderer := Renderer{}
	req, err := renderer.NewRequest(http.MethodPost, "/", td1)
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
