package postgres

import (
	"fmt"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestPostgres(t *testing.T) {
	godotenv.Load("../../.env")
	c := GetPostgresConfig()
	actual := c.GetPostgresConnectionInfo()
	expected := "host=localhost port=5432 user=user password=yoko1893 dbname=gacha sslmode=disable"

	if assert.Equal(t, expected, actual) {
		fmt.Printf("%v", "good")
	} else {
		t.Errorf("c.GetPostgresConnectionInfo() = %v, want %v", c.GetPostgresConnectionInfo(), expected)
	}
}
