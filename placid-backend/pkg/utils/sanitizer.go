package utils

import (
	"placid-backend/internal/models"
	placiderror "placid-backend/internal/placid_error"
	"strings"
)

func SanitizeAuthRequest(user *models.User) error {

	if len(user.Email) < 3 {
		return placiderror.ErrInvalidEmail
	}

	if !strings.Contains(user.Email, "@") {
		return placiderror.ErrInvalidEmail

	}

	if len(user.Password) < 6 {
		return placiderror.ErrInsecurePassword
	}

	return nil

}
