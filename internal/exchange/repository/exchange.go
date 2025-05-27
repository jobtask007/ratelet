package repository

import (
	"errors"
	"ratelet/internal/exchange/domain"
)

var (
	errNotFound = errors.New("not found")
)

type Exchange struct {
	data []CryptoCurrencyRate
}

func NewExchange() *Exchange {
	return &Exchange{
		data: []CryptoCurrencyRate{
			{"BEER", 18, 0.00002461},
			{"FLOKI", 18, 0.0001428},
			{"GATE", 18, 6.87},
			{"USDT", 6, 0.999},
			{"WBTC", 8, 57037.22},
		},
	}
}

func (r *Exchange) GetCryptoRate(symbol string) (domain.CryptoCurrencyRate, error) {
	for _, d := range r.data {
		if symbol == d.Symbol {
			return d.MapToDomain(), nil
		}
	}

	return domain.CryptoCurrencyRate{}, errNotFound
}
