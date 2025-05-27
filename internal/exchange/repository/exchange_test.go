package repository

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"ratelet/internal/exchange/domain"
	"testing"
)

func TestExchange_GetCryptoRate(t *testing.T) {
	ex := NewExchange()

	tests := []struct {
		name         string
		symbol       string
		expectedRate domain.CryptoCurrencyRate
		expectedErr  error
	}{
		{
			name:   "Success",
			symbol: "USDT",
			expectedRate: domain.CryptoCurrencyRate{
				Symbol:        "USDT",
				DecimalPlaces: 6,
				RateToUSD:     0.999,
			},
			expectedErr: nil,
		},
		{
			name:         "Fails on symbol not found",
			symbol:       "UNKNOWN",
			expectedRate: domain.CryptoCurrencyRate{},
			expectedErr:  errNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// when
			rate, err := ex.GetCryptoRate(tt.symbol)

			// then
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expectedRate, rate)
		})
	}
}
