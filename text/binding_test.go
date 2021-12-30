package text

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Binder_BindFromRequest(t *testing.T) {
	td1 := "test"
	var td2 string

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(td1)))
	binder := Binder{}

	binder.BindFromRequest(req, &td2)
	if td1 != td2 {
		t.Fatalf("%s != %s", td1, td2)
	}
}

func Test_Binder_BindFromResponse(t *testing.T) {
	td1 := "test"
	var td2 string

	res := &http.Response{
		Header:        make(http.Header),
		Body:          ioutil.NopCloser(bytes.NewBuffer([]byte(td1))),
		ContentLength: int64(len(td1)),
	}

	binder := Binder{}
	binder.BindFromResponse(res, &td2)
	if td1 != td2 {
		t.Fatalf("%s != %s", td1, td2)
	}
}
