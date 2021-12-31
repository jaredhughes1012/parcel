package json

import (
	"encoding/json"
	"errors"
	"mime"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jaredhughes1012/parcel/httperror"
)

func testJsonResponse(t *testing.T, w *httptest.ResponseRecorder, status int, target interface{}) {
	if w.Result().StatusCode != status {
		t.Errorf("Status code invalid: %d != %d", w.Result().StatusCode, status)
	}

	if ct, _, err := mime.ParseMediaType(w.Header().Get("Content-Type")); err != nil {
		t.Fatal(err)
	} else if ct != "application/json" {
		t.Errorf("Content type invalid: %s != %s", ct, "application/json")
	} else if err = json.NewDecoder(w.Body).Decode(target); err != nil {
		t.Fatal(err)
	}
}

func testJsonRequest(t *testing.T, req *http.Request, target interface{}) {
	if ct, _, err := mime.ParseMediaType(req.Header.Get("Content-Type")); err != nil {
		t.Fatal(err)
	} else if ct != "application/json" {
		t.Errorf("Content type invalid: %s != %s", ct, "application/json")
	} else if err = json.NewDecoder(req.Body).Decode(target); err != nil {
		t.Fatal(err)
	}
}

func Test_Renderer_RenderResponse(t *testing.T) {
	type testData struct {
		Field string `json:"field"`
	}

	td1 := testData{Field: "test"}
	var td2 testData

	w := httptest.NewRecorder()
	renderer := Renderer{}

	renderer.RenderResponse(w, http.StatusCreated, &td1)
	testJsonResponse(t, w, http.StatusCreated, &td2)
	if td1.Field != td2.Field {
		t.Errorf("Rendering failed: %s != %s", td1.Field, td2.Field)
	}
}

func Test_Renderer_NewRequest(t *testing.T) {
	type testData struct {
		Field string `json:"field"`
	}

	td1 := testData{Field: "test"}

	renderer := Renderer{}
	req, err := renderer.NewRequest(http.MethodPost, "/", &td1)
	if err != nil {
		t.Fatal(err)
	}

	var td2 testData
	testJsonRequest(t, req, &td2)
	if td1.Field != td2.Field {
		t.Errorf("Rendering failed: %s != %s", td1.Field, td2.Field)
	}
}

func Test_Renderer_RenderError_HttpError(t *testing.T) {
	td1 := httperror.New(http.StatusConflict, "test")
	var td2 httperror.HttpError

	w := httptest.NewRecorder()
	renderer := Renderer{}

	renderer.RenderErrorResponse(w, td1)
	testJsonResponse(t, w, http.StatusConflict, &td2)

	if td1.Error() != td2.Error() {
		t.Errorf("Rendering failed: %s != %s", td1.Error(), td2.Error())
	}
}

func Test_Renderer_RenderError_NonHttpError(t *testing.T) {
	td1 := errors.New("test")
	var td2 httperror.HttpError

	w := httptest.NewRecorder()
	renderer := Renderer{}

	renderer.RenderErrorResponse(w, td1)

	testJsonResponse(t, w, http.StatusInternalServerError, &td2)
	if !strings.Contains(td2.Error(), td1.Error()) {
		t.Errorf("Rendering failed: %s != %s", td1.Error(), td2.Error())
	}
}
