package endpoints

import (
	"context"
	"time"

	"github.com/denniszl/wallet_flexing/internal/flex"
	"github.com/pkg/errors"
)

// GetTransactionsRequest is a request for getting transactions
type GetTransactionsRequest struct {
	GetAll bool
	From   time.Time
	To     time.Time
}

// GetTransactionsResponse is a response for getting transactions
type GetTransactionsResponse struct {
	DateTime string  `json:"datetime"`
	Amount   float64 `json:"amount"`
}

// maybe use this if we can't use RFC3339
// const outputTimestampFormat = "2006-01-02T15:04:05+07:00"

func convertTypes(t []flex.Transaction) []GetTransactionsResponse {
	converted := []GetTransactionsResponse{}
	for _, transaction := range t {
		converted = append(converted, GetTransactionsResponse{
			DateTime: transaction.Timestamp.Format(time.RFC3339),
			Amount:   transaction.Amount,
		})
	}
	return converted
}

// GetTransactions gets transactions
func (e endpoints) GetTransactions(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(GetTransactionsRequest)
	var transactions []GetTransactionsResponse
	if req.GetAll {
		allTransactions, err := e.service.GetAllAmounts(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get all amounts")
		}

		transactions = convertTypes(allTransactions)
	} else {
		someTransactions, err := e.service.GetAmountsFromTime(ctx, req.From, req.To)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get amounts")
		}

		transactions = convertTypes(someTransactions)
	}

	return transactions, nil
}
