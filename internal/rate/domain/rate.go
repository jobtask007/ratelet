package domain

type Rates struct {
	Rates []Rate `json:"rates"`
}
type Rate struct {
	From string  `json:"from"`
	To   string  `json:"to"`
	Rate float64 `json:"rate"`
}
