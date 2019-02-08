package config

import (
	"github.com/apex/log"
	"github.com/caarlos0/env"
)

// Config of the app
type Config struct {
	Account  string `env:"ACCOUNT" envDefault:"454"`
	Password string `env:"PASSWORD" envDefault:"454"`
	Identity string `env:"IDENTITY" envDefault:"51eec6356c615d3edf39d497c137d75b"`
}

// MustGet returns the config
func MustGet() Config {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.WithError(err).Fatal("failed to load config")
	}
	return cfg
}
