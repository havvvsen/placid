package config

import (
	"testing"
	"github.com/gofiber/fiber/v3"
)

func TestConfigureCors(t *testing.T) {
	app := fiber.New()
	allowedOrigins := []string{"http://localhost:3000"}

	// Ensure no panic
	ConfigureCors(app, allowedOrigins)

	// In a real scenario, we might want to test the response headers,
	// but here we just ensure the configuration can be attached to the app.
	if app == nil {
		t.Fatalf("App should not be nil")
	}
}
