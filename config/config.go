package config

import (
	"github.com/apex/log"
	"github.com/caarlos0/env"
)

// Config of the app
type Config struct {
	Port                    string   `env:"PORT" envDefault:"8080"`
	Account                 string   `env:"ACCOUNT" envDefault:"123"`
	Password                string   `env:"PASSWORD" envDefault:"123"`
	Identity                string   `env:"IDENTITY" envDefault:"fd11vvosaxnf8gskg39j"`
	AhgoraURL               string   `env:"AHGORA_URL" envDefault:"https://www.ahgora.com.br"`
	LogLevel                string   `env:"LOG_LEVEL" envDefault:"DEBUG"`
	AhgoraMockServerEnabled bool     `env:"AHGORA_MOCK_SERVER_ENABLED" envDefault:"false"`
	CronTimes               []string `env:"CRON_TIMES" envSeparator:"," envDefault:"08:00,11:50,13:20,17:50"`
}

// MustGet returns the config
func MustGet() Config {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.WithError(err).Fatal("failed to load config")
	}
	return cfg
}
