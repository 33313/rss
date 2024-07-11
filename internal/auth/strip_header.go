package auth

import (
	"errors"
	"net/http"
	"strings"
)

var ErrNoAuthHeader = errors.New("Missing authorization header")
var ErrBadAuthHeader = errors.New("Malformed authorization header")

func StripHeader(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")
	if header == "" {
		return "", ErrNoAuthHeader
	}
	split := strings.Split(header, " ")
	if len(split) != 2 {
		return "", ErrBadAuthHeader
	}
	return split[1], nil
}
