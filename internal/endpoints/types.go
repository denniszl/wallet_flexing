package endpoints

import (
	"context"

	"github.com/denniszl/wallet_flexing/internal/flex"
)

// Transaction is a btc transaction
type Transaction struct {
	DateTime string  `json:"datetime"`
	Amount   float64 `json:"amount"`
}

// Endpoints is an interface for rateplan endpoints.
type Endpoints interface {
	GetTransactions(context.Context, interface{}) (interface{}, error)

	PostTransaction(context.Context, interface{}) (interface{}, error)
}

type endpoints struct {
	service flex.Service
}

// NewEndpoints returns an endpoints struct
func NewEndpoints(s flex.Service) Endpoints {
	return endpoints{
		service: s,
	}
}
