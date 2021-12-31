package text

import (
	"errors"
	"io"
	"mime"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jaredhughes1012/parcel/httperror"
)

func testTextResponse(t *testing.T, w *httptest.ResponseRecorder, status int) string {
	if w.Result().StatusCode != status {
		t.Errorf("Status code invalid: %d != %d", w.Result().StatusCode, status)
	}

	if ct, _, err := mime.ParseMediaType(w.Header().Get("Content-Type")); err != nil {
		t.Fatal(err)
	} else if ct != "text/plain" {
		t.Errorf("Content type invalid: %s != %s", ct, "text/plain")
	}

	b, err := io.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	}

	return string(b)
}

func testJsonRequest(t *testing.T, req *http.Request) string {
	if ct, _, err := mime.ParseMediaType(req.Header.Get("Content-Type")); err != nil {
		t.Fatal(err)
	} else if ct != "text/plain" {
		t.Errorf("Content type invalid: %s != %s", ct, "text/plain")
	}

	b, err := io.ReadAll(req.Body)
	if err != nil {
		t.Fatal(err)
	}

	return string(b)
}

func Test_Renderer_RenderResponse(t *testing.T) {
	td1 := "test"

	w := httptest.NewRecorder()
	renderer := Renderer{}

	renderer.RenderResponse(w, http.StatusCreated, td1)
	s := testTextResponse(t, w, http.StatusCreated)
	if s != td1 {
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

	s := testJsonRequest(t, req)
	if s != td1 {
		t.Errorf("Rendered value invalid: %s != %s", td1, s)
	}
}

func Test_Renderer_RenderError_HttpError(t *testing.T) {
	he := httperror.New(http.StatusConflict, "test").(*httperror.HttpError)

	w := httptest.NewRecorder()
	renderer := Renderer{}

	renderer.RenderErrorResponse(w, he)
	s := testTextResponse(t, w, http.StatusConflict)
	if s != he.Message {
		t.Errorf("Rendered value invalid: %s != %s", he.Message, s)
	}
}

func Test_Renderer_RenderError_NonHttpError(t *testing.T) {
	e := errors.New("test")

	w := httptest.NewRecorder()
	renderer := Renderer{}

	renderer.RenderErrorResponse(w, e)
	s := testTextResponse(t, w, http.StatusInternalServerError)
	if s != e.Error() {
		t.Errorf("Rendered value invalid: %s != %s", e.Error(), s)
	}
}
