package disk

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/denniszl/wallet_flexing/internal/repository"
	"github.com/pkg/errors"
)

func (d *diskRepo) GetTransactions(ctx context.Context) ([]repository.Transaction, error) {
	d.lock.RLock()
	defer d.lock.RUnlock()
	transactions := []repository.Transaction{}
	// read from transactions.csv
	file, err := os.Open("transactions.csv")
	if err != nil {
		return nil, errors.Wrap(err, "failed to read transactions.csv")
	}

	defer file.Close()

	r := csv.NewReader(file)
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, errors.Wrap(err, "failed to read csv line")
		}
		timestamp, err := time.Parse(time.RFC3339, record[0])
		if err != nil {
			return nil, errors.Wrap(err, "incorrect format for timstamp")
		}

		amount, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return nil, errors.Wrap(err, "must have some numeric value for amount")
		}

		transactions = append(transactions, repository.Transaction{
			Timestamp: timestamp,
			Amount:    amount,
		})
	}

	sort.Slice(transactions, func(i, j int) bool { return transactions[i].Timestamp.Before(transactions[j].Timestamp) })

	return transactions, nil
}

func (d *diskRepo) SaveTransaction(ctx context.Context, timestamp time.Time, amount float64) error {
	d.lock.Lock()
	defer d.lock.Unlock()
	// write to transactions.csv
	f, err := os.OpenFile("transactions.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return errors.Wrap(err, "failed to open csv")
	}
	defer f.Close()
	w := csv.NewWriter(f)
	err = w.Write([]string{timestamp.Format(time.RFC3339), fmt.Sprintf("%f", d.amountHeld+amount)})
	if err != nil {
		return errors.Wrap(err, "failed to write to csv")
	}
	w.Flush()
	d.amountHeld += amount
	return nil
}
