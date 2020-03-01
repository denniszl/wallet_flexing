package flex

import (
	"context"
	"time"

	"github.com/denniszl/wallet_flexing/internal/repository"
)

// Transaction represents a transaction of being given BTC
type Transaction struct {
	Timestamp time.Time
	Amount    float64
}

// Service for flexing
type Service interface {
	GetAllAmounts(ctx context.Context) ([]Transaction, error)
	GetAmountsFromTime(ctx context.Context, to time.Time, from time.Time) ([]Transaction, error)

	SaveAmount(ctx context.Context, timestamp time.Time, amount float64) error
}

type flexClient struct {
	repo repository.Repository
}

func NewService(r repository.Repository) Service {
	return &flexClient{
		repo: r,
	}
}
