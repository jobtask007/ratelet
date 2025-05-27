//go:build unit

package config

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name               string
		envVars            map[string]string
		expectedPort       uint
		expectedLogLevel   string
		expectedOxrAPIHost string
		expectedOxrAppID   string
		expectedDevMode    bool
		expectedErr        error
	}{
		{
			name: "Success both environment variables set",
			envVars: map[string]string{
				"PORT":                         "8181",
				"LOG_LEVEL":                    "DEBUG",
				"OPEN_EXCHANGE_RATES_API_HOST": "https://openexchangerates.org/api/v2",
				"OPEN_EXCHANGE_RATES_APP_ID":   "621726",
				"DEV_MODE":                     "true",
			},
			expectedPort:       8181,
			expectedLogLevel:   "DEBUG",
			expectedOxrAPIHost: "https://openexchangerates.org/api/v2",
			expectedOxrAppID:   "621726",
			expectedDevMode:    true,
			expectedErr:        nil,
		},
		{
			name: "Success with default environment variables",
			envVars: map[string]string{
				"OPEN_EXCHANGE_RATES_APP_ID": "621726",
			},
			expectedPort:       8080,
			expectedLogLevel:   "INFO",
			expectedOxrAPIHost: "https://openexchangerates.org/api",
			expectedOxrAppID:   "621726",
			expectedDevMode:    false,
			expectedErr:        nil,
		},
		{
			name: "Fail on invalid PORT value",
			envVars: map[string]string{
				"OPEN_EXCHANGE_RATES_APP_ID": "621726",
				"PORT":                       "-10",
			},
			expectedErr: errors.New("parsing env vars: env: parse error on field \"Port\" of type \"uint\": strconv.ParseUint: parsing \"-10\": invalid syntax"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			for key, value := range tt.envVars {
				t.Setenv(key, value)
			}

			// when
			config, err := NewConfig()

			// then
			if tt.expectedErr != nil {
				assert.Nil(t, config)
				assert.EqualError(t, err, tt.expectedErr.Error())

				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expectedPort, config.Port, tt.expectedPort, config.Port)
			assert.Equal(t, tt.expectedLogLevel, config.LogLevel)
			assert.Equal(t, tt.expectedOxrAPIHost, config.OpenExchangeRatesAPIHost)
			assert.Equal(t, tt.expectedOxrAppID, config.OpenExchangeRatesAppID)
			assert.Equal(t, tt.expectedDevMode, config.DevMode)
		})
	}
}
