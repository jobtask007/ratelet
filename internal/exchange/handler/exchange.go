package handler

import (
	"log/slog"
	"net/http"
	"strconv"

	"ratelet/internal/exchange/dto"

	"github.com/gin-gonic/gin"
)

type Service interface {
	Exchange(from, to string, amount float64) (string, error)
}

type Exchange struct {
	svc Service
}

func New(s Service) *Exchange {
	return &Exchange{
		svc: s,
	}
}

func (e Exchange) Exchange(c *gin.Context) {
	fromParam := c.Query("from")
	toParam := c.Query("to")
	amountParam := c.Query("amount")

	if fromParam == "" || toParam == "" || amountParam == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	amount, err := strconv.ParseFloat(amountParam, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := e.svc.Exchange(fromParam, toParam, amount)
	if err != nil {
		slog.Error("Failed to exchange cryptocurrencies",
			"from", fromParam, "to", toParam, "amount", amountParam, "err", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	response := dto.ExchangeResponse{
		From:   fromParam,
		To:     toParam,
		Amount: result,
	}

	c.JSON(http.StatusOK, response)
}
