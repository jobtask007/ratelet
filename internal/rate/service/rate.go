package service

import (
	"fmt"
	"log/slog"
	"math"

	"ratelet/internal/oxr"
	"ratelet/internal/rate/domain"
)

type Service struct {
	oxrClient OxrClient
}

type OxrClient interface {
	GetRates(currencies []string) (oxr.RatesResponse, error)
}

func New(oxrClient OxrClient) *Service {
	return &Service{oxrClient: oxrClient}
}

func (s *Service) GetRates(currencies []string) (domain.Rates, error) {
	oxrRates, err := s.oxrClient.GetRates(currencies)
	if err != nil {
		return domain.Rates{}, fmt.Errorf("getting oxr rates: %w", err)
	}

	base := oxrRates.Base

	var rates []string

	for c := range oxrRates.Rates {
		rates = append(rates, c)
	}

	var result domain.Rates

	for i := 0; i < len(rates); i++ {
		for j := i + 1; j < len(rates); j++ {
			from := rates[i]
			to := rates[j]

			rateFromTo := getRate(from, to, base, oxrRates.Rates)
			rateToFrom := getRate(to, from, base, oxrRates.Rates)

			result.Rates = append(result.Rates, domain.Rate{
				From: from,
				To:   to,
				Rate: math.Round(rateFromTo*100) / 100,
			})

			result.Rates = append(result.Rates, domain.Rate{
				From: to,
				To:   from,
				Rate: math.Round(rateToFrom*100) / 100,
			})
		}
	}

	return result, nil
}

func getRate(from, to, base string, rates map[string]float64) float64 {
	if from == base {
		return rates[to]
	}

	if to == base {
		r := rates[from]
		if r == 0 {
			slog.Warn("Rate from base is zero", "currency", from)
			return 0
		}

		return 1 / r
	}

	rateFrom := rates[from]
	rateTo := rates[to]
	if rateFrom == 0 || rateTo == 0 {
		return 0
	}

	return rateTo / rateFrom
}
