package utils

import (
	"placid-backend/internal/models"
	placiderror "placid-backend/internal/placid_error"
	"strings"
)

func SanitizeAuthRequest(request *models.RegisterRequest) error {

	if len(request.Email) < 3 {
		return placiderror.ErrInvalidEmail
	}

	if !strings.Contains(request.Email, "@") {
		return placiderror.ErrInvalidEmail

	}

	if len(request.Password) < 6 {
		return placiderror.ErrInsecurePassword
	}

	return nil

}
