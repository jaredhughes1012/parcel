package json

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Binder_BindFromRequest(t *testing.T) {
	type testData struct {
		Field string `json:"field"`
	}

	td1 := testData{Field: "test"}
	var td2 testData

	b, _ := json.Marshal(&td1)
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(b))
	binder := Binder{}

	binder.BindFromRequest(req, &td2)
	if td1.Field != td2.Field {
		t.Fatalf("%s != %s", td1.Field, td2.Field)
	}
}

func Test_Binder_BindFromResponse(t *testing.T) {
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

	binder := Binder{}
	binder.BindFromResponse(res, &td2)
	if td1.Field != td2.Field {
		t.Fatalf("%s != %s", td1.Field, td2.Field)
	}
}
