package http

import (
	"errors"
	"net/http"
)

var (
	// ErrBadRouting is returned when an expected path variable is missing.
	// It always indicates programmer error.
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
	// ErrBadContext is returned when an expected context key is missing.
	// It always indicates programmer error.
	ErrBadContext = errors.New("expected key to exist in Context but not found (programmer error")
	// ErrNotFound is returned when route not found
	ErrNotFound = errors.New("not Found")
)

// mapHTTPError will map endpoint and transport errors to http errors
func mapHTTPError(err error) (status int, found bool) {
	found = true
	switch err {
	case ErrNotFound:
		status = http.StatusNotFound
	default:
		found = false
	}
	return
}

// ErrorResponse holds error responses
type ErrorResponse struct {
	Errors []Error `json:"errors"`
}

// Error holds error messages
type Error struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}
