package config

import (
	"fmt"

	"github.com/caarlos0/env/v8"
)

// Config struct holds all the configuration for the application
type Config struct {
	DBUrl    string `env:"DB_URL" envDefault:"root:pass@/kvadolibrarydb"`
	Port     int    `env:"PORT" envDefault:"9000"`
	Local    bool   `env:"LOCAL" envDefault:"false"`
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
}

// NewConfig function creates a new Config struct by parsing the environment variables
func NewConfig() (*Config, error) {
	var c Config
	if err := env.Parse(&c); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	return &c, nil
}
