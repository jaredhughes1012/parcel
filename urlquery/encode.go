package urlquery

import (
	"net/url"
	"strconv"
)

// Adds a boolean value to the give url values
func EncodeBool(q *url.Values, name string, value bool) {
	if value {
		q.Set(name, "true")
	} else {
		q.Set(name, "false")
	}
}

// Adds an integer value to the give url values
func EncodeInt(q *url.Values, name string, value int) {
	v := strconv.Itoa(value)
	q.Set(name, v)
}
