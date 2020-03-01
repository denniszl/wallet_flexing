package http

import (
	"context"
	"net/http"

	"github.com/denniszl/wallet_flexing/internal/endpoints"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// MakeHTTPHandler initializes a go-kit http service
func MakeHTTPHandler(serviceEndpoints endpoints.Endpoints) http.Handler {
	r := mux.NewRouter()

	// maybe one day
	options := []httptransport.ServerOption{}
	options = append(
		options,
		httptransport.ServerErrorEncoder(makeErrorEncoder()),
	)

	notFound := func(_ context.Context, _ interface{}) (interface{}, error) {
		return nil, ErrNotFound
	}
	r.NotFoundHandler = httptransport.NewServer(
		notFound,
		decodeRequest,
		encodeResponse,
		options...,
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

	// GET /transactions
	r.Methods("GET").Path("/transactions").Handler(
		httptransport.NewServer(
			serviceEndpoints.GetTransactions,
			decodeGetTransactions,
			encodeResponse,
			options...,
		),
	).Name("transactions")

	// POST /transactions
	r.Methods("POST").Path("/transactions").Handler(
		httptransport.NewServer(
			serviceEndpoints.PostTransaction,
			decodePostTransaction,
			encodeResponse,
			options...,
		),
	).Name("transactions")

	return r
}
