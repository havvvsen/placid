package config

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func ConfigureCors(app *fiber.App, allowedCORSOrigins []string) {
	app.Use(
		cors.New(
			cors.Config{
				AllowOrigins: allowedCORSOrigins,
				AllowMethods: []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
				AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
			},
		),
	)
}
