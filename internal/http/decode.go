package http

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

// Decoder type to be for things like wrappers -- similar to endpoint.Endpoint.
type Decoder func(c context.Context, r *http.Request) (request interface{}, err error)

// DecoderWrapper is the return type of functions that will be used for chaining
// Decoder code.
type DecoderWrapper func(Decoder) httptransport.DecodeRequestFunc

// decodeRequest is a default decoder when we don't need anything special
func decodeRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return nil, nil
}
