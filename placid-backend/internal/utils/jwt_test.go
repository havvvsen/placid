package utils

import (
	"testing"
	"github.com/golang-jwt/jwt/v5"
)

func TestGenerateJWTToken(t *testing.T) {
	secret := "mysecretkey"
	email := "test@example.com"
	isAdmin := true

	tokenString, err := GenerateJWTToken(secret, email, isAdmin)
	if err != nil {
		t.Fatalf("GenerateJWTToken returned error: %v", err)
	}
	if tokenString == "" {
		t.Fatalf("Expected token string, got empty string")
	}

	// Verify the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	})
	if err != nil {
		t.Fatalf("Failed to parse token: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["email"] != email {
			t.Errorf("Expected email %v, got %v", email, claims["email"])
		}
		if claims["is_admin"] != isAdmin {
			t.Errorf("Expected is_admin %v, got %v", isAdmin, claims["is_admin"])
		}
	} else {
		t.Fatalf("Invalid token or claims")
	}
}
