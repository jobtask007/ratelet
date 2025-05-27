package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Port                     uint   `env:"PORT" envDefault:"8080"`
	LogLevel                 string `env:"LOG_LEVEL" envDefault:"INFO"`
	OpenExchangeRatesAPIHost string `env:"OPEN_EXCHANGE_RATES_API_HOST" envDefault:"https://openexchangerates.org/api"`
	OpenExchangeRatesAppID   string `env:"OPEN_EXCHANGE_RATES_APP_ID,required"`
	DevMode                  bool   `env:"DEV_MODE" envDefault:"false"`
}

func NewConfig() (*Config, error) {
	c := &Config{}

	err := env.Parse(c)
	if err != nil {
		return nil, fmt.Errorf("parsing env vars: %w", err)
	}

	return c, nil
}
