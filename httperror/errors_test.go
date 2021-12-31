package httperror

import (
	"errors"
	"net/http"
	"strings"
	"testing"
)

func Test_Wrap(t *testing.T) {
	err := errors.New("test")
	wrapped := Wrap(http.StatusConflict, err)

	he, ok := wrapped.(*HttpError)
	if !ok {
		t.Fatal("Error is not an HttpError")
	}

	if he.Status != http.StatusConflict {
		t.Errorf("Status failed: %d != %d", he.Status, http.StatusConflict)
	}
	if !strings.Contains(he.Error(), err.Error()) {
		t.Errorf("Incorrect error message: %s", he.Error())
	}
}

func Test_Convert_HttpError(t *testing.T) {
	err := &HttpError{
		Status:  http.StatusConflict,
		wrapped: errors.New("test"),
	}

	he := Convert(error(err))

	if he.Status != err.Status {
		t.Errorf("Status failed: %d != %d", he.Status, err.Status)
	}
	if he.Error() != err.Error() {
		t.Errorf("Error failed: %s != %s", he.Error(), err.Error())
	}
}

func Test_Convert_NotHttpError(t *testing.T) {
	err := errors.New("test")

	he := Convert(error(err))

	if he.Status != http.StatusInternalServerError {
		t.Errorf("Status failed: %d != %d", he.Status, http.StatusInternalServerError)
	}
	if he.wrapped.Error() != err.Error() {
		t.Errorf("Error wrapping failed: %s != %s", he.wrapped.Error(), err.Error())
	}
}
