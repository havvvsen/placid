package middlewares

import (
	"fmt"
	"log/slog"
	placiderror "placid-backend/internal/placid_error"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func AdminMiddleware(logger *slog.Logger) fiber.Handler {
	return func(c fiber.Ctx) error {
		claims := c.Locals("claims").(jwt.MapClaims)
		email := claims["email"]
		role := claims["role"]

		if role != "admin" {
			logger.Info(fmt.Sprintf("Unauthorized admin action attempt by: %s", email))
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": placiderror.ErrUnauthorizedOperation.Error()})
		}

		return c.Next()
	}

}
