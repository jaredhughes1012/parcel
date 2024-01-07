package check

import (
	"net/http"
	"testing"

	"github.com/jaredhughes1012/parcel/httperror"
	"github.com/stretchr/testify/assert"
)

func Test_Mime(t *testing.T) {
	cases := []struct {
		contentType string
		header      string
		errStatus   int
	}{
		{
			contentType: "application/json",
			header:      "application/json",
		},
		{
			contentType: "application/json",
			header:      "application/xml",
			errStatus:   http.StatusUnsupportedMediaType,
		},
		{
			contentType: "application/json",
			header:      "",
			errStatus:   http.StatusBadRequest,
		},
	}

	for _, c := range cases {
		err := Mime(c.header, c.contentType)
		if c.errStatus == 0 {
			assert.NoError(t, err)
		} else {
			status, e := httperror.Unwrap(err)
			assert.Error(t, e)
			assert.Equal(t, c.errStatus, status)
		}
	}
}
