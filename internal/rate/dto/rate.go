package dto

import "ratelet/internal/rate/domain"

type RateResponse struct {
	From string  `json:"from"`
	To   string  `json:"to"`
	Rate float64 `json:"rate"`
}

func MapFromDomain(rates domain.Rates) []RateResponse {
	result := make([]RateResponse, len(rates.Rates))

	for i, r := range rates.Rates {
		result[i] = RateResponse{
			From: r.From,
			To:   r.To,
			Rate: r.Rate,
		}
	}

	return result
}
