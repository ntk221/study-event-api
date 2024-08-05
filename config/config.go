package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	DBHost     string        `env:"DB_HOST" envDefault:"localhost"`
	DBPort     int           `env:"DB_PORT" envDefault:"5432"`
	DBUser     string        `env:"DB_USER" envDefault:"testuser"`
	DBPassword string        `env:"DB_PASSWORD" envDefault:"testpass"`
	DBName     string        `env:"DB_NAME" envDefault:"study_event_api_test"`
	DBSSLMode  string        `env:"DB_SSLMODE" envDefault:"disable"`
	ServerPort int           `env:"SERVER_PORT" envDefault:"8080"`
	Timeout    time.Duration `env:"TIMEOUT" envDefault:"30s"`
}

func Load() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to parse environment variables: %w", err)
	}
	return cfg, nil
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName, c.DBSSLMode)
}
