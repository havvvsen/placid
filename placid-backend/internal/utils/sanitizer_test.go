package utils

import (
	"testing"
	"placid-backend/internal/models"
	placiderror "placid-backend/internal/placid_error"
)

func TestSanitizeAuthRequest(t *testing.T) {
	tests := []struct {
		name     string
		request  *models.RegisterRequest
		expected error
	}{
		{
			name: "Valid request",
			request: &models.RegisterRequest{
				Email:    "test@example.com",
				Password: "password123",
			},
			expected: nil,
		},
		{
			name: "Email too short",
			request: &models.RegisterRequest{
				Email:    "a@",
				Password: "password123",
			},
			expected: placiderror.ErrInvalidEmail,
		},
		{
			name: "Email missing @",
			request: &models.RegisterRequest{
				Email:    "testexample.com",
				Password: "password123",
			},
			expected: placiderror.ErrInvalidEmail,
		},
		{
			name: "Password too short",
			request: &models.RegisterRequest{
				Email:    "test@example.com",
				Password: "pass",
			},
			expected: placiderror.ErrInsecurePassword,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := SanitizeAuthRequest(tt.request)
			if err != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, err)
			}
		})
	}
}
