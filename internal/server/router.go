package server

import (
	exch "ratelet/internal/exchange/handler"
	rate "ratelet/internal/rate/handler"
)

func (s *Server) RegisterRoutes(rates *rate.Rates, exch *exch.Exchange) {
	s.engine.GET("/rates", rates.GetRates)

	s.engine.GET("/exchange", exch.Exchange)
}
