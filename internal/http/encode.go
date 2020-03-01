package http

import (
	"context"
	"encoding/json"
	"net/http"
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
