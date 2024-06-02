package amount_test

import (
	"fmt"
	"testing"

	"github.com/HLerman/test/internal/libs/amount"
	"github.com/stretchr/testify/require"
)

func TestConvertToCents(t *testing.T) {
	type Result struct {
		Err   error
		Cents int
	}
	tests := map[float64]Result{
		123.45: {
			Err:   nil,
			Cents: 12345,
		},
		123.456: {
			Err:   amount.ErrInvalidAmount,
			Cents: 0,
		},
		12.3: {
			Err:   nil,
			Cents: 1230,
		},
		12.34: {
			Err:   nil,
			Cents: 1234,
		},
		12.00000001: {
			Err:   amount.ErrInvalidAmount,
			Cents: 0,
		},
		12.001: {
			Err:   amount.ErrInvalidAmount,
			Cents: 0,
		},
		9.7: {
			Err:   nil,
			Cents: 970,
		},
		9.333: {
			Err:   amount.ErrInvalidAmount,
			Cents: 0,
		},
		9.33: {
			Err:   nil,
			Cents: 933,
		},
		9.99: {
			Err:   nil,
			Cents: 999,
		},
	}

	for montant := range tests {
		t.Run("case "+fmt.Sprintf("%f", montant), func(t *testing.T) {
			cents, err := amount.ConvertToCents(montant)
			require.Equal(t, tests[montant].Err, err)
			require.Equal(t, tests[montant].Cents, cents)
		})
	}
}
