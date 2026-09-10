package config

import (
	"fmt"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	PostgresDSN string `env:"POSTGRES_DSN,required"`
	HttpPort    string `env:"HTTP_PORT,required"`
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to parse .env: %w", err)
	}
	return cfg, nil
}
