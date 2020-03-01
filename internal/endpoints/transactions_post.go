package endpoints

import (
	"context"
	"errors"
	"time"
)

// PostTransactionRequest is a request to post a transaction
type PostTransactionRequest struct {
	Timestamp string  `json:"datetime"`
	Amount    float64 `json:"amount"`
}

var (
	// ErrInvalidTimestamp is when the timestamp is an invalid format
	ErrInvalidTimestamp = errors.New("timestamp must be RFC3339")
	// ErrTimestampFromFuture is when timestamp given in is after time.Now()
	ErrTimestampFromFuture = errors.New("transaction is from the future")
	// ErrInvalidAmount is when the amount is <= 0
	ErrInvalidAmount = errors.New("must send an amount that's greater than 0")
)

// PostTransaction posts transactions
func (e endpoints) PostTransaction(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(PostTransactionRequest)

	if req.Amount <= 0.0 {
		return nil, ErrInvalidAmount
	}
	timestamp, err := time.Parse(time.RFC3339, req.Timestamp)
	if err != nil {
		return nil, ErrInvalidTimestamp
	}

	if timestamp.After(time.Now()) {
		return nil, ErrTimestampFromFuture
	}

	e.service.SaveAmount(ctx, timestamp, req.Amount)
	response := map[string]bool{}
	response["status"] = true
	return response, nil
}
