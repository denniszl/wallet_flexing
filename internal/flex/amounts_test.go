package flex

import (
	"context"
	"testing"
	"time"

	"github.com/denniszl/wallet_flexing/internal/repository"
	"github.com/stretchr/testify/require"
)

/*
	GetTransactions(context.Context) ([]Transaction, error)
	SaveTransaction(context.Context, time.Time, float64) error

*/

type mockRepo struct {
	repository.Repository
	amts []repository.Transaction
}

func (m *mockRepo) GetTransactions(_ context.Context) ([]repository.Transaction, error) {
	return m.amts, nil
}

func (m *mockRepo) SaveTransaction(_ context.Context, t time.Time, amt float64) error {
	m.amts = append(m.amts, repository.Transaction{
		Timestamp: t,
		Amount:    amt,
	})

	return nil
}

func TestGetAllAmounts(t *testing.T) {
	now := time.Now()
	oneHourBefore := now.Add(-1 * time.Hour)
	halfanHourBefore := now.Add(-30 * time.Minute)
	twoHoursBefore := now.Add(-2 * time.Hour)
	testCases := []struct {
		testName      string
		amts          []repository.Transaction
		expectedAmts  []Transaction
		expectedError bool
	}{
		{
			testName:      "repo returns nothing",
			amts:          []repository.Transaction{},
			expectedAmts:  []Transaction{},
			expectedError: false,
		},
		{
			testName: "repo returns a single record",
			amts: []repository.Transaction{
				{
					Timestamp: oneHourBefore,
					Amount:    10.0,
				},
			},
			expectedAmts: []Transaction{
				{
					Timestamp: oneHourBefore,
					Amount:    10.0,
				},
			},
			expectedError: false,
		},
		{
			testName: "repo returns multiple records",
			amts: []repository.Transaction{
				{
					Timestamp: twoHoursBefore,
					Amount:    8.0,
				},
				{
					Timestamp: oneHourBefore,
					Amount:    10.0,
				},
				{
					Timestamp: halfanHourBefore,
					Amount:    12.0,
				},
			},
			expectedAmts: []Transaction{
				{
					Timestamp: twoHoursBefore,
					Amount:    8.0,
				},
				{
					Timestamp: oneHourBefore,
					Amount:    10.0,
				},
				{
					Timestamp: halfanHourBefore,
					Amount:    12.0,
				},
			},
			expectedError: false,
		},
	}

	for _, tc := range testCases {
		r := &mockRepo{}
		r.amts = tc.amts
		s := NewService(r)
		t.Run(tc.testName, func(t *testing.T) {
			amts, err := s.GetAllAmounts(context.Background())
			require.Equal(t, tc.expectedError, err != nil)
			require.Equal(t, tc.expectedAmts, amts)
		})
	}
}

func TestGetAmountsFromTime(t *testing.T) {
	now := time.Now()
	oneHourBefore := now.Add(-1 * time.Hour)
	halfanHourBefore := now.Add(-30 * time.Minute)
	twoHoursBefore := now.Add(-2 * time.Hour)
	testCases := []struct {
		testName      string
		amts          []repository.Transaction
		expectedAmts  []Transaction
		from          time.Time
		to            time.Time
		expectedError bool
	}{
		{
			testName:      "repo returns nothing",
			amts:          []repository.Transaction{},
			expectedAmts:  []Transaction{},
			expectedError: false,
			from:          now.Add(-3 * time.Hour),
			to:            now.Add(-2 * time.Hour),
		},
		{
			testName:      "repo errors when times are misaligned",
			amts:          []repository.Transaction{},
			expectedAmts:  nil,
			expectedError: true,
			to:            now.Add(-3 * time.Hour),
			from:          now.Add(-2 * time.Hour),
		},
		{
			testName: "repo returns a single record",
			amts: []repository.Transaction{
				{
					Timestamp: oneHourBefore,
					Amount:    10.0,
				},
			},
			expectedAmts: []Transaction{
				{
					Timestamp: oneHourBefore,
					Amount:    10.0,
				},
			},
			expectedError: false,
			from:          now.Add(-1 * time.Hour),
			to:            now,
		},
		{
			testName: "repo returns multiple records",
			amts: []repository.Transaction{
				{
					Timestamp: twoHoursBefore,
					Amount:    8.0,
				},
				{
					Timestamp: oneHourBefore,
					Amount:    10.0,
				},
				{
					Timestamp: halfanHourBefore,
					Amount:    12.0,
				},
			},
			expectedAmts: []Transaction{
				{
					Timestamp: twoHoursBefore,
					Amount:    8.0,
				},
				{
					Timestamp: oneHourBefore,
					Amount:    10.0,
				},
				{
					Timestamp: halfanHourBefore,
					Amount:    12.0,
				},
			},
			expectedError: false,
			from:          now.Add(-2 * time.Hour),
			to:            now,
		},
		{
			testName: "repo returns multiple records",
			amts: []repository.Transaction{
				{
					Timestamp: twoHoursBefore,
					Amount:    8.0,
				},
				{
					Timestamp: oneHourBefore,
					Amount:    10.0,
				},
				{
					Timestamp: halfanHourBefore,
					Amount:    12.0,
				},
			},
			expectedAmts: []Transaction{
				{
					Timestamp: oneHourBefore,
					Amount:    10.0,
				},
				{
					Timestamp: halfanHourBefore,
					Amount:    12.0,
				},
			},
			expectedError: false,
			from:          now.Add(-1 * time.Hour),
			to:            now,
		},
	}

	for _, tc := range testCases {
		r := &mockRepo{}
		r.amts = tc.amts
		s := NewService(r)
		t.Run(tc.testName, func(t *testing.T) {
			amts, err := s.GetAmountsFromTime(context.Background(), tc.from, tc.to)
			require.Equal(t, tc.expectedError, err != nil)
			require.Equal(t, tc.expectedAmts, amts)
		})
	}
}
