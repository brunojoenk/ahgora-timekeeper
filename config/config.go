package config

import (
	"github.com/apex/log"
	"github.com/caarlos0/env"
)

// Config of the app
type Config struct {
	Account  string `env:"ACCOUNT"`
	Password string `env:"PASSWORD"`
	Identity string `env:"IDENTITY"`
}

// MustGet returns the config
func MustGet() Config {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.WithError(err).Fatal("failed to load config")
	}
	return cfg
}
