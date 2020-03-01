package repository

import (
	"context"
	"time"
)

type Transaction struct {
	Timestamp time.Time
	Amount    float64
}

// Repository describes what the interface can do
type Repository interface {
	GetTransactions(context.Context) ([]Transaction, error)
	SaveTransaction(context.Context, time.Time, float64) error
}
