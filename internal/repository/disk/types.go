package disk

import (
	"context"
	"os"
	"sync"

	"github.com/denniszl/wallet_flexing/internal/repository"
	"github.com/pkg/errors"
)

// implementation of the repository that writes to a file on disk.
type diskRepo struct {
	// lock so we don't mess up the file we're writing to
	lock       sync.RWMutex
	amountHeld float64
}

// NewRepository returns a disk repository. It will return
func NewRepository() (repository.Repository, error) {
	if _, err := os.Stat("transactions.csv"); os.IsNotExist(err) {
		f, err := os.Create("transactions.csv")
		if err != nil {
			return nil, err
		}
		defer f.Close()
	}
	// initialize amount held:
	r := &diskRepo{
		lock: sync.RWMutex{},
	}

	transactions, err := r.GetTransactions(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize repo")
	}
	total := 0.0
	if len(transactions) > 0 {
		total = transactions[len(transactions)-1].Amount
	}

	r.amountHeld = total

	return r, nil
}
