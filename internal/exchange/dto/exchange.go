package dto

type ExchangeResponse struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount string `json:"amount"`
}
