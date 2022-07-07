package parcel

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jaredhughes1012/parcel/httperror"
)

func Test_BindFromRequest_Json(t *testing.T) {
	type testData struct {
		Field string `json:"field"`
	}

	td1 := testData{Field: "test"}
	var td2 testData

	b, _ := json.Marshal(&td1)
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")

	err := BindFromRequest(req, &td2)
	if err != nil {
		t.Fatal()
	} else if td1.Field != td2.Field {
		t.Fatalf("%s != %s", td1.Field, td2.Field)
	}
}

func Test_BindFromRequest_UnsupportedObjectType(t *testing.T) {
	type testData struct {
		Field string `json:"field"`
	}

	td1 := testData{Field: "test"}
	var td2 testData

	b, _ := json.Marshal(&td1)
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "evil/data")

	hErr := BindFromRequest(req, &td2)
	status, err := httperror.Unwrap(hErr)

	if err == nil {
		t.Fatal("Non http error returned")
	} else if status != http.StatusUnsupportedMediaType {
		t.Fatalf("Invalid http status code %d", status)
	}
}

func Test_BindFromRequest_NoMediaType(t *testing.T) {
	type testData struct {
		Field string `json:"field"`
	}

	td1 := testData{Field: "test"}
	var td2 testData

	b, _ := json.Marshal(&td1)
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(b))

	hErr := BindFromRequest(req, &td2)
	status, err := httperror.Unwrap(hErr)

	if err == nil {
		t.Fatal("Non http error returned")
	} else if status != http.StatusBadRequest {
		t.Fatalf("Invalid http status code %d", status)
	}
}

func Test_BindFromRequest_Text(t *testing.T) {
	td1 := "test"
	var td2 string

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(td1)))
	req.Header.Set("Content-Type", "text/plain")

	err := BindFromRequest(req, &td2)
	if err != nil {
		t.Fatal()
	} else if td1 != td2 {
		t.Fatalf("%s != %s", td1, td2)
	}
}

func Test_BindFromResponse_Json(t *testing.T) {
	type testData struct {
		Field string `json:"field"`
	}

	td1 := testData{Field: "test"}
	var td2 testData

	b, _ := json.Marshal(&td1)
	res := &http.Response{
		Header:        make(http.Header),
		Body:          ioutil.NopCloser(bytes.NewBuffer(b)),
		ContentLength: int64(len(b)),
	}
	res.Header.Set("Content-Type", "application/json")

	err := BindFromResponse(res, &td2)
	if err != nil {
		t.Fatal()
	} else if td1.Field != td2.Field {
		t.Fatalf("%s != %s", td1.Field, td2.Field)
	}
}

func Test_BindFromResponse_UnsupportedObjectType(t *testing.T) {
	type testData struct {
		Field string `json:"field"`
	}

	td1 := testData{Field: "test"}
	var td2 testData

	b, _ := json.Marshal(&td1)
	res := &http.Response{
		Header:        make(http.Header),
		Body:          ioutil.NopCloser(bytes.NewBuffer(b)),
		ContentLength: int64(len(b)),
	}
	res.Header.Set("Content-Type", "evil/data")

	hErr := BindFromResponse(res, &td2)
	status, err := httperror.Unwrap(hErr)

	if err == nil {
		t.Fatal("Non http error returned")
	} else if status != http.StatusUnsupportedMediaType {
		t.Fatalf("Invalid http status code %d", status)
	}
}

func Test_BindFromResponse_NoMediaType(t *testing.T) {
	type testData struct {
		Field string `json:"field"`
	}

	td1 := testData{Field: "test"}
	var td2 testData

	b, _ := json.Marshal(&td1)
	res := &http.Response{
		Header:        make(http.Header),
		Body:          ioutil.NopCloser(bytes.NewBuffer(b)),
		ContentLength: int64(len(b)),
	}

	hErr := BindFromResponse(res, &td2)
	status, err := httperror.Unwrap(hErr)

	if err == nil {
		t.Fatal("Non http error returned")
	} else if status != http.StatusBadRequest {
		t.Fatalf("Invalid http status code %d", status)
	}
}

func Test_BindFromResponse_Text(t *testing.T) {
	td1 := "test"
	var td2 string

	res := &http.Response{
		Header:        make(http.Header),
		Body:          ioutil.NopCloser(bytes.NewBuffer([]byte(td1))),
		ContentLength: int64(len(td1)),
	}
	res.Header.Set("Content-Type", "text/plain")

	err := BindFromResponse(res, &td2)
	if err != nil {
		t.Fatal()
	} else if td1 != td2 {
		t.Fatalf("%s != %s", td1, td2)
	}
}
