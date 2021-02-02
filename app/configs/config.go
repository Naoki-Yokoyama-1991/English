package configs

import (
	"os"

	"github.com/naoki/gacha/app/postgres"
)

const (
	prod = "production"
)

// Config object
type Config struct {
	Postgres postgres.PostgresConfig `json:"postgres"`
	Env      string                  `env:"ENV"`
	Host     string                  `env:"APP_HOST"`
	Port     string                  `env:"APP_PORT"`
}

// IsPros Checks if env is production
func (c Config) IsProd() bool {
	return c.Env == prod
}

// GetConfig gets all config for the application
func GetConfig() Config {
	return Config{
		Env:      os.Getenv("ENV"),
		Postgres: postgres.GetPostgresConfig(),
		Host:     os.Getenv("APP_HOST"),
		Port:     os.Getenv("APP_PORT"),
	}
}
