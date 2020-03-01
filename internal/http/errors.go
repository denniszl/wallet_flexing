package http

import (
	"errors"
	"net/http"

	"github.com/denniszl/wallet_flexing/internal/endpoints"
)

var (
	// ErrBadRouting is returned when an expected path variable is missing.
	// It always indicates programmer error.
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
	// ErrBadContext is returned when an expected context key is missing.
	// It always indicates programmer error.
	ErrBadContext = errors.New("expected key to exist in Context but not found (programmer error")
	// ErrNotFound is returned when route not found
	ErrNotFound = errors.New("not found")
	// ErrInvalidTimestampsFormat uses an invalid timestamp format
	ErrInvalidTimestampsFormat = errors.New("must use RFC3339 for timestamps")
	// ErrInvalidTimestamps is when the start date is after the end date.
	ErrInvalidTimestamps = errors.New("the start datetime must be before the end datetime")
	// ErrMissingTimestamp means you're missing either a start or end timestamp
	ErrMissingTimestamp = errors.New("must provide both a start_date_time and an end_date_time")
	// ErrInvalidBody is when the body is invalid
	ErrInvalidBody = errors.New("invalid body for request")
)

// mapHTTPError will map endpoint and transport errors to http errors
func mapHTTPError(err error) (status int, found bool) {
	found = true
	switch err {
	case ErrNotFound:
		status = http.StatusNotFound
	case ErrInvalidTimestamps, ErrInvalidTimestampsFormat, ErrMissingTimestamp, endpoints.ErrTimestampFromFuture, endpoints.ErrInvalidTimestamp, endpoints.ErrInvalidAmount, ErrInvalidBody:
		status = http.StatusBadRequest
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
	Type    string `json:"type,omitempty"`
	Message string `json:"message"`
}
