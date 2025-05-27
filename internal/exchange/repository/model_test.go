package repository

import (
	"testing"

	"ratelet/internal/exchange/domain"

	"github.com/stretchr/testify/assert"
)

func TestCryptoCurrencyRate_MapToDomain(t *testing.T) {
	// given
	c := CryptoCurrencyRate{
		Symbol:        "USDT",
		DecimalPlaces: 6,
		RateToUSD:     0.999,
	}

	// when
	got := c.MapToDomain()

	// then
	expected := domain.CryptoCurrencyRate{
		Symbol:        "USDT",
		DecimalPlaces: 6,
		RateToUSD:     0.999,
	}

	assert.Equal(t, expected, got)
}
