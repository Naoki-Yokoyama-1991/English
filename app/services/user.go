package services

import (
	"html"
	"strings"
	"time"

	"github.com/naoki/gacha/app/models"
)

func Prepare(user *models.User) {
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
}
