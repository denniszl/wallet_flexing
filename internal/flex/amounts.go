package flex

import (
	"context"
	"fmt"
	"time"

	"github.com/denniszl/wallet_flexing/internal/repository"
)

func transformTransactions(t []repository.Transaction) []Transaction {
	transactions := []Transaction{}
	for _, t := range t {
		transactions = append(transactions, Transaction{
			Timestamp: t.Timestamp,
			Amount:    t.Amount,
		})
	}

	return transactions
}

// GetAllAmounts gets everything stored in the repository and outputs it in utc
func (f *flexClient) GetAllAmounts(ctx context.Context) ([]Transaction, error) {
	transactions, err := f.repo.GetTransactions(ctx)
	if err != nil {
		return nil, err
	}
	return transformTransactions(transactions), nil
}

func transactionInInterval(transaction Transaction, lo time.Time, hi time.Time) bool {
	ts := transaction.Timestamp
	return (ts.Equal(lo) || ts.Equal(hi)) || (ts.After(lo) && ts.Before(hi))
}

// GetAmountsFromTime gets everything from the repository and filters out things from to -> from
// probably better to improve this in the future by making that happen in the repository layer.
// times are converted to utc
func (f *flexClient) GetAmountsFromTime(ctx context.Context, from time.Time, to time.Time) ([]Transaction, error) {
	if to.Before(from) {
		return nil, fmt.Errorf("endDatetime must be after startDatetime")
	}

	transactions, err := f.repo.GetTransactions(ctx)
	if err != nil {
		return nil, err
	}

	transactionsInInterval := []Transaction{}

	transformedTransactions := transformTransactions(transactions)

	for _, t := range transformedTransactions {
		if transactionInInterval(t, from, to) {
			transactionsInInterval = append(transactionsInInterval, t)
		}
	}

	return transactionsInInterval, nil
}

// SaveAmount saves an amount to the repository with a timestamp of now.
func (f *flexClient) SaveAmount(ctx context.Context, timestamp time.Time, amount float64) error {
	return f.repo.SaveTransaction(ctx, timestamp, amount)
}
