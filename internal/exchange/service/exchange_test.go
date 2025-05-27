package service

import (
	"errors"
	"testing"

	"ratelet/internal/exchange/domain"
	"ratelet/internal/exchange/service/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestService_Exchange(t *testing.T) {
	tests := []struct {
		name           string
		from           string
		to             string
		amount         float64
		mockSetup      func(c *mocks.MockRepository)
		expectedResult string
		expectedErr    error
	}{
		{
			name:   "Success",
			from:   "WBTC",
			to:     "USDT",
			amount: 1,
			mockSetup: func(c *mocks.MockRepository) {
				c.On("GetCryptoRate", "WBTC").
					Return(domain.CryptoCurrencyRate{
						Symbol:        "WBTC",
						DecimalPlaces: 8,
						RateToUSD:     57037.22,
					}, nil)
				c.On("GetCryptoRate", "USDT").
					Return(domain.CryptoCurrencyRate{
						Symbol:        "USDT",
						DecimalPlaces: 6,
						RateToUSD:     0.999,
					}, nil)
			},
			expectedResult: "57094.314314",
			expectedErr:    nil,
		},
		{
			name:   "Fails on getting 'from' rate",
			from:   "WBTC",
			to:     "USDT",
			amount: 1,
			mockSetup: func(c *mocks.MockRepository) {
				c.On("GetCryptoRate", mock.Anything).
					Return(domain.CryptoCurrencyRate{}, errors.New("not found"))
			},
			expectedResult: "",
			expectedErr:    errors.New("not found"),
		}, {
			name:   "Fails on getting 'to' rate",
			from:   "WBTC",
			to:     "USDT",
			amount: 1,
			mockSetup: func(c *mocks.MockRepository) {
				c.On("GetCryptoRate", "WBTC").
					Return(domain.CryptoCurrencyRate{
						Symbol:        "WBTC",
						DecimalPlaces: 8,
						RateToUSD:     57037.22,
					}, nil)
				c.On("GetCryptoRate", "USDT").
					Return(domain.CryptoCurrencyRate{}, errors.New("not found"))
			},
			expectedResult: "",
			expectedErr:    errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			mockRepo := new(mocks.MockRepository)
			svc := New(mockRepo)

			if tt.mockSetup != nil {
				tt.mockSetup(mockRepo)
			}

			// when
			amount, err := svc.Exchange(tt.from, tt.to, tt.amount)

			// then
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expectedResult, amount)

			mockRepo.AssertExpectations(t)
		})
	}
}
