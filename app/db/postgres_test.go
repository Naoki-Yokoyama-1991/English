package db

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestPostgres(t *testing.T) {
	godotenv.Load("../../.env")
	c := GetPostgresConfig()
	actual := c.GetPostgresConnectionInfo()
	expected := "host=postgres port=5432 user=admin password=admin dbname=admin sslmode=disable"

	assert.Equal(t, expected, actual)
}
