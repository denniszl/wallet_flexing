package http

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// MakeHTTPHandler initializes a go-kit http service
func MakeHTTPHandler() http.Handler {
	r := mux.NewRouter()

	// maybe one day
	options := []httptransport.ServerOption{}

	notFound := func(_ context.Context, _ interface{}) (interface{}, error) {
		return nil, ErrNotFound
	}
	r.NotFoundHandler = httptransport.NewServer(
		notFound,
		decodeRequest,
		encodeResponse,
		options...,
	)

	// Favicon not found
	r.Methods("GET").Path("/favicon.ico").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		},
	)

	// GET /_healthcheck
	healthcheck := func(_ context.Context, _ interface{}) (interface{}, error) {
		response := map[string]bool{}
		response["Status"] = true
		return response, nil
	}
	r.Methods("GET").Path("/_healthcheck").Handler(
		httptransport.NewServer(
			healthcheck,
			decodeRequest,
			encodeResponse,
			options...,
		),
	).Name("healthcheck")

	// GET /_heartbeat
	r.Methods("GET").Path("/_heartbeat").Handler(
		httptransport.NewServer(
			healthcheck,
			decodeRequest,
			encodeResponse,
			options...,
		),
	).Name("heartbeat")

	return r
}
