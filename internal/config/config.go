package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	DataPath string `env:"DATA_PATH" envDefault:"data/apk_fingerprints.json"`
	Host     string `env:"HOST" envDefault:"0.0.0.0"`
	Port     int    `env:"PORT" envDefault:"8080"`
	AuthKey  string `env:"AUTH_KEY,required"`
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("Can not load env: %w", err)
	}
	config, err := env.ParseAs[Config]()

	if err != nil {
		return nil, fmt.Errorf("Can not load config: %w", err)
	}

	return &config, nil
}
