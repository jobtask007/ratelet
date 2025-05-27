package service

import (
	"errors"
	"testing"

	"ratelet/internal/oxr"
	"ratelet/internal/rate/domain"
	"ratelet/internal/rate/service/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_GetRates(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		currencies     []string
		mockSetup      func(c *mocks.MockOxrClient)
		expectedResult domain.Rates
		expectedErr    error
	}{
		{
			name:       "Success with base currency",
			currencies: []string{"USD", "PLN"},
			mockSetup: func(c *mocks.MockOxrClient) {
				c.On("GetRates", []string{"USD", "PLN"}).
					Return(oxr.RatesResponse{
						Base: "USD",
						Rates: map[string]float64{
							"USD": 1.0,
							"PLN": 3.846721,
						},
					}, nil)
			},
			expectedResult: domain.Rates{
				Rates: []domain.Rate{
					{From: "USD", To: "PLN", Rate: 3.85},
					{From: "PLN", To: "USD", Rate: 0.26},
				},
			},
			expectedErr: nil,
		},
		{
			name:       "Success without base currency",
			currencies: []string{"ZAR", "PLN"},
			mockSetup: func(c *mocks.MockOxrClient) {
				c.On("GetRates", []string{"ZAR", "PLN"}).
					Return(oxr.RatesResponse{
						Base: "USD",
						Rates: map[string]float64{
							"ZAR": 17.89421,
							"PLN": 3.846721,
						},
					}, nil)
			},
			expectedResult: domain.Rates{
				Rates: []domain.Rate{
					{From: "ZAR", To: "PLN", Rate: 0.21},
					{From: "PLN", To: "ZAR", Rate: 4.65},
				},
			},
			expectedErr: nil,
		},
		{
			name:       "Success when base currency rate is zero",
			currencies: []string{"USD", "PLN"},
			mockSetup: func(c *mocks.MockOxrClient) {
				c.On("GetRates", []string{"USD", "PLN"}).
					Return(oxr.RatesResponse{
						Base: "USD",
						Rates: map[string]float64{
							"USD": 1.0,
							"PLN": 0,
						},
					}, nil)
			},
			expectedResult: domain.Rates{
				Rates: []domain.Rate{
					{From: "USD", To: "PLN", Rate: 0},
					{From: "PLN", To: "USD", Rate: 0},
				},
			},
			expectedErr: nil,
		}, {
			name:       "Success when non-base currency rate is zero",
			currencies: []string{"ZAR", "PLN"},
			mockSetup: func(c *mocks.MockOxrClient) {
				c.On("GetRates", []string{"ZAR", "PLN"}).
					Return(oxr.RatesResponse{
						Base: "USD",
						Rates: map[string]float64{
							"ZAR": 17.89421,
							"PLN": 0,
						},
					}, nil)
			},
			expectedResult: domain.Rates{
				Rates: []domain.Rate{
					{From: "ZAR", To: "PLN", Rate: 0},
					{From: "PLN", To: "ZAR", Rate: 0},
				},
			},
			expectedErr: nil,
		},
		{
			name:       "When oxrClient returns error",
			currencies: []string{"USD", "PLN"},
			mockSetup: func(c *mocks.MockOxrClient) {
				c.On("GetRates", []string{"USD", "PLN"}).
					Return(oxr.RatesResponse{}, errors.New("oxr client error"))
			},
			expectedResult: domain.Rates{},
			expectedErr:    errors.New("getting oxr rates: oxr client error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			mockOxrClient := new(mocks.MockOxrClient)
			svc := New(mockOxrClient)

			if tt.mockSetup != nil {
				tt.mockSetup(mockOxrClient)
			}

			// when
			rates, err := svc.GetRates(tt.currencies)

			// then
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
				return
			}

			require.NoError(t, err)
			assert.ElementsMatch(t, tt.expectedResult.Rates, rates.Rates)

			mockOxrClient.AssertExpectations(t)
		})
	}
}
