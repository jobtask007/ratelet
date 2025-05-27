package handler

import (
	"log/slog"
	"net/http"
	"strings"

	"ratelet/internal/rate/domain"
	"ratelet/internal/rate/dto"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetRates(currencies []string) (domain.Rates, error)
}

type Rates struct {
	svc Service
}

func New(s Service) *Rates {
	return &Rates{
		svc: s,
	}
}

func (r Rates) GetRates(c *gin.Context) {
	currenciesParam := c.Query("currencies")

	if currenciesParam == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	currencies := strings.Split(currenciesParam, ",")

	if len(currencies) < 2 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	rates, err := r.svc.GetRates(currencies)
	if err != nil {
		slog.Error("Failed to get rates", "currencies", currencies, "err", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	response := dto.MapFromDomain(rates)

	c.JSON(http.StatusOK, response)
}
