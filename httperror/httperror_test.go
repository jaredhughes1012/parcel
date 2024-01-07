package httperror

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	message := "this is a test message"
	result := New(http.StatusNotFound, message)

	assert.Equal(t, http.StatusNotFound, result.Code)
	assert.True(t, strings.Contains(result.Error(), message))
	assert.True(t, strings.Contains(result.Error(), fmt.Sprintf("%d", http.StatusNotFound)))
}

func Test_Wrap(t *testing.T) {
	err := errors.New("this is a test message")
	result := Wrap(http.StatusConflict, err)

	assert.Equal(t, http.StatusConflict, result.Code)
	assert.True(t, strings.Contains(result.Error(), err.Error()))
	assert.True(t, strings.Contains(result.Error(), fmt.Sprintf("%d", http.StatusConflict)))
}

func Test_Fmt(t *testing.T) {
	message := "this is a test message"

	result := Fmt(http.StatusBadGateway, "%s", message)

	assert.Equal(t, http.StatusBadGateway, result.Code)
	assert.True(t, strings.Contains(result.Error(), message))
	assert.True(t, strings.Contains(result.Error(), fmt.Sprintf("%d", http.StatusBadGateway)))
}

func Test_Unwrap(t *testing.T) {
	cases := []struct {
    name string
		err  error
		code int
	}{
		{
			err:  nil,
			code: -1,
		},
		{
			err:  errors.New("this is a test message"),
			code: http.StatusInternalServerError,
		},
		{
			err:  &HttpError{Code: http.StatusConflict, Wrapped: errors.New("this is a test message")},
			code: http.StatusConflict,
		},
	}

  for _, c := range cases {
    t.Run(c.name, func(t *testing.T) {
      code, err := Unwrap(c.err)
      assert.Equal(t, c.code, code)

      if h, ok := c.err.(*HttpError); ok {
        assert.Equal(t, h.Wrapped, err)
      } else {
        assert.Equal(t, c.err, err)
      }
    })
  }
}
