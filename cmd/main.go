package main

import (
	"log/slog"
	"os"

	"ratelet/internal/config"
	exchhnd "ratelet/internal/exchange/handler"
	"ratelet/internal/exchange/repository"
	exchsvc "ratelet/internal/exchange/service"
	"ratelet/internal/oxr"
	ratehnd "ratelet/internal/rate/handler"
	ratesvc "ratelet/internal/rate/service"
	"ratelet/internal/server"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error("Failed to create config", "err", err)
		os.Exit(1)
	}

	initLogger(cfg.LogLevel)

	oxrClient := oxr.NewClient(cfg.OpenExchangeRatesAPIHost, cfg.OpenExchangeRatesAppID)

	ratesSvc := ratesvc.New(oxrClient)
	ratesHandler := ratehnd.New(ratesSvc)

	repo := repository.NewExchange()
	exchSvc := exchsvc.New(repo)
	exchangeHandler := exchhnd.New(exchSvc)

	srv := server.New(cfg.Port, cfg.DevMode)
	srv.RegisterRoutes(ratesHandler, exchangeHandler)
	srv.Run()
}

func initLogger(level string) {
	logLevel := &slog.LevelVar{}

	err := logLevel.UnmarshalText([]byte(level))
	if err != nil {
		slog.Error("Failed to parse log level", "err", err)
		os.Exit(1)
	}

	opts := &slog.HandlerOptions{
		Level: logLevel,
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	slog.SetDefault(logger)
}
