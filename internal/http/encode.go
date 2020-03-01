package http

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/prometheus/common/log"
)

// encodeResponse is a default encoder we use when we just need to marshal JSON
func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err, ok := response.(error); ok {
		// determine status code
		httpStatusCode, found := mapHTTPError(err)
		if !found {
			httpStatusCode = http.StatusBadRequest
		}
		w.WriteHeader(httpStatusCode)
		// write marshaled error response
		return json.NewEncoder(w).Encode(ErrorResponse{
			Errors: []Error{
				Error{
					Message: err.Error(),
				},
			},
		})
	}
	return json.NewEncoder(w).Encode(response)
}

// makeErrorEncoder creates an encoder used by ServerErrorEncoder
func makeErrorEncoder() httptransport.ErrorEncoder {
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		if err == nil {
			panic("ErrorEncoder called with nil error")
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		// determine HTTP status code
		httpStatusCode, found := mapHTTPError(err)
		if !found {
			httpStatusCode = http.StatusInternalServerError
		}
		w.WriteHeader(httpStatusCode)
		// for 5XX errors, return generic error message
		errorMessage := err.Error()
		if httpStatusCode >= 500 && httpStatusCode <= 599 {
			errorMessage = http.StatusText(httpStatusCode)
		}

		if httpStatusCode != http.StatusNotFound {
			// log error
			log.Error(
				"encoded error: " + err.Error(),
			)
		} else {
			log.Info("not found")
		}
		// write marshaled error response
		json.NewEncoder(w).Encode(ErrorResponse{
			Errors: []Error{
				Error{
					Message: errorMessage,
				},
			},
		})
	}
}
