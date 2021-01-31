package configs

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestPostgres(t *testing.T) {
	t.Log("start")

	godotenv.Load("../../.env")
	c := GetConfig()

	t.Run("log1", func(t *testing.T) {
		t.Log("log1")
		actual := c.Postgres.GetPostgresConnectionInfo()
		expected := "host=postgres port=5432 user=admin password=admin dbname=admin sslmode=disable"
		assert.Equal(t, expected, actual)
	})

	t.Run("log2", func(t *testing.T) {
		t.Log("log2")
		actual := c.Postgres.Dialect()
		expected := "postgres"
		assert.Equal(t, expected, actual)
	})

	t.Log("finish")
}
