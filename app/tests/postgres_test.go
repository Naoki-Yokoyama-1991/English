package tests

import (
	"fmt"
	"testing"

	"github.com/joho/godotenv"
	"github.com/naoki/english/app/postgres"
	"github.com/stretchr/testify/assert"
)

func TestPostgres(t *testing.T) {
	godotenv.Load("../../.env")
	c := postgres.GetPostgresConfig()
	actual := c.GetPostgresConnectionInfo()
	expected := "host=postgres port=5432 user=admin password=admin dbname=admin sslmode=disable"

	if assert.Equal(t, expected, actual) {
		fmt.Printf("%v", "good")
	} else {
		t.Errorf("c.GetPostgresConnectionInfo() = %v, want %v", c.GetPostgresConnectionInfo(), expected)
	}
}
