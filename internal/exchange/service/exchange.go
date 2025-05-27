package service

import (
	"log/slog"

	"ratelet/internal/exchange/domain"

	"github.com/shopspring/decimal"
)

type Service struct {
	repo Repository
}

type Repository interface {
	GetCryptoRate(symbol string) (domain.CryptoCurrencyRate, error)
}

func New(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) Exchange(from, to string, amount float64) (string, error) {
	rateFrom, err := s.repo.GetCryptoRate(from)
	if err != nil {
		slog.Error("Failed to get crypto rate", "symbol", from, "err", err)
		return "", err
	}

	rateTo, err := s.repo.GetCryptoRate(to)
	if err != nil {
		slog.Error("Failed to get crypto rate", "symbol", to, "err", err)
		return "", err
	}

	decAmount := decimal.NewFromFloatWithExponent(amount, -30)
	decRateFrom := decimal.NewFromFloatWithExponent(rateFrom.RateToUSD, -30)
	decRateTo := decimal.NewFromFloatWithExponent(rateTo.RateToUSD, -30)

	amountInUSD := decAmount.Mul(decRateFrom)
	result := amountInUSD.Div(decRateTo)
	rounded := result.Round(int32(rateTo.DecimalPlaces))

	return rounded.String(), nil
}
