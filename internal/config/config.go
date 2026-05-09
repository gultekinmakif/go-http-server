package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Port string
	Env  string // dev or prod
}

func getEnv(key string, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func Load() (*Config, error) {
	cfg := &Config{
		Port: getEnv("PORT", "3000"),
		Env:  getEnv("ENV", "dev"),
	}

	if _, err := strconv.Atoi(cfg.Port); err != nil {
		return nil, fmt.Errorf("invalid PORT %q: must be a number", cfg.Port)
	}

	if cfg.Env != "dev" && cfg.Env != "prod" {
		return nil, fmt.Errorf("invalid ENV %q: must be a 'dev' or 'prod'", cfg.Env)
	}

	return cfg, nil
}
