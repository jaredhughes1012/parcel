package urlquery

import (
	"fmt"
	"net/url"
	"strconv"
)

// Decodes a boolean value from the url's query params. Will return an error if
// the value is not set
func DecodeBool(u url.URL, name string) (bool, error) {
	v := u.Query().Get(name)
	if v == "true" {
		return true, nil
	} else if v == "false" {
		return false, nil
	} else {
		return false, fmt.Errorf("could not decode param %v=%v into bool", name, v)
	}
}

// Decodes a boolean value from the url's query params. Will return the default value
// if the value is not set
func DecodeBoolDefault(u url.URL, name string, defaultVal bool) bool {
	v, err := DecodeBool(u, name)
	if err != nil {
		return defaultVal
	} else {
		return v
	}
}

// Decodes an integer value from the url's query params. Will return an error if
// the value is not set
func DecodeInt(u url.URL, name string) (int, error) {
	return strconv.Atoi(u.Query().Get(name))
}

// Decodes an integer value from the url's query params. Will return the default value
// if the value is not set
func DecodeIntDefault(u url.URL, name string, defaultVal int) int {
	v, err := DecodeInt(u, name)
	if err != nil {
		return defaultVal
	} else {
		return v
	}
}

// Gets a string value from the url's query params. Will return the default value
// if the value is not set
func DecodeStringDefault(u url.URL, name string, defaultVal string) string {
	v := u.Query().Get(name)
	if v == "" {
		return defaultVal
	} else {
		return v
	}
}
