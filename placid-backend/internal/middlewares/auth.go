package middlewares

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(logger *slog.Logger, jwtSecretKey string) fiber.Handler {
	return func(c fiber.Ctx) error {
		if !c.HasHeader("Authorization") {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing authorization header"})
		}

		authorizationHeader := c.Get("Authorization", "")
		if authorizationHeader == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Blank authorization header"})
		}

		authorizationHeaderItems := strings.Split(authorizationHeader, " ")

		if len(authorizationHeaderItems) != 2 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Blank authorization header"})
		}

		if authorizationHeaderItems[0] != "Bearer" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Only Bearer token is supported"})
		}

		tokenString := authorizationHeaderItems[len(authorizationHeaderItems)-1]
		tokenString = strings.Trim(tokenString, " ")
		tokenString = strings.Trim(tokenString, "\n")

		token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(t *jwt.Token) (any, error) {
			return []byte(jwtSecretKey), nil
		})

		if err != nil || !token.Valid {
			logger.Error(fmt.Sprintf("Failed to parse JWT token string: %s", err.Error()))
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		c.Locals("claims", token.Claims.(jwt.MapClaims))

		return c.Next()
	}

}
