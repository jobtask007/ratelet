//go:build unit

package dto

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"ratelet/internal/rate/domain"
)

func TestMapFromDomain(t *testing.T) {
	// given
	rates := domain.Rates{
		Rates: []domain.Rate{
			{From: "USD", To: "PLN", Rate: 3.85},
			{From: "PLN", To: "USD", Rate: 0.26},
		},
	}

	// when
	ratesResponse := MapFromDomain(rates)

	// then
	expected := []RateResponse{
		{From: "USD", To: "PLN", Rate: 3.85},
		{From: "PLN", To: "USD", Rate: 0.26},
	}

	assert.ElementsMatch(t, expected, ratesResponse)
}
