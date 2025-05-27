package repository

import "ratelet/internal/exchange/domain"

type CryptoCurrencyRate struct {
	Symbol        string
	DecimalPlaces int
	RateToUSD     float64
}

func (c CryptoCurrencyRate) MapToDomain() domain.CryptoCurrencyRate {
	return domain.CryptoCurrencyRate{
		Symbol:        c.Symbol,
		DecimalPlaces: c.DecimalPlaces,
		RateToUSD:     c.RateToUSD,
	}
}
