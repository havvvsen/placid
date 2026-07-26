package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTToken(jwtSecretKey string, email string, isAdmin bool) (string, error) {
	var tokenString string

	// Create jwt token valid for 30 days
	// TODO: implement refresh token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":    email,
		"is_admin": isAdmin,
		"exp":      time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtSecretKey))

	if err != nil {
		return tokenString, err
	}

	return tokenString, nil
}
