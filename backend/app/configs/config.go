package configs

import (
	"os"

	"github.com/naoki/gacha/app/db"
)

const (
	prod = "production"
)

// Config object
type Config struct {
	Env      string            `env:"ENV"`
	Host     string            `env:"APP_HOST"`
	Port     string            `env:"APP_PORT"`
	Postgres db.PostgresConfig `json:"postgres"`
}

// IsPros Checks if env is production
func (c Config) IsProd() bool {
	return c.Env == prod
}

// GetConfig gets all config for the application
func GetConfig() Config {
	return Config{
		Env:      os.Getenv("ENV"),
		Postgres: db.GetPostgresConfig(),
		Host:     os.Getenv("APP_HOST"),
		Port:     os.Getenv("APP_PORT"),
	}
}
