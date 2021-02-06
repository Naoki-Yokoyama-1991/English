package validate

import (
	"errors"
	"strings"

	"github.com/naoki/task/app/models"
)

func Validate(user *models.User, action string) map[string]string {
	var errorMessages = make(map[string]string)
	var err error
	switch strings.ToLower(action) {
	case "update":
		if "userEmail" == "" {
			err = errors.New("Required Email")
			errorMessages["Required_email"] = err.Error()
		}
	default:
		if user.Username == "" {
			err = errors.New("Required Password")
			errorMessages["Required_username"] = err.Error()
		}
	}
	return errorMessages
}
