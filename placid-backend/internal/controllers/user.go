package controllers

import (
	"errors"
	"fmt"
	"placid-backend/internal/models"
	"placid-backend/internal/placid"
	placiderror "placid-backend/internal/placid_error"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
)

func ControllerUser(server *placid.Server) fiber.Handler {
	return func(c fiber.Ctx) error {
		ctx := c.Context()
		var request models.UserRequest

		if c.HasBody() {
			if err := c.Bind().Body(&request); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": placiderror.ErrBadRequest.Error()})
			}

			user, err := server.Queries.GetUser(ctx, pgtype.Text{
				Valid:  true,
				String: request.Email,
			})

			if err != nil {
				if errors.Is(err, pgx.ErrNoRows) {
					return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": placiderror.ErrUnauthorizedOperation.Error()})
				}
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": placiderror.ErrInternalServerError.Error()})
			}

			return c.Status(fiber.StatusOK).JSON(
				fiber.Map{
					"message": models.UserResponse{
						Uuid:      user.Uuid.String(),
						Email:     user.Email.String,
						IsAdmin:   user.IsAdmin.Bool,
						IsPremium: user.IsPremium.Bool,
						CreatedAt: user.CreatedAt.Time.String(),
					},
				},
			)

		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": placiderror.ErrBadRequest.Error()})
	}
}

func ControllerDeleteAccount(server *placid.Server) fiber.Handler {
	return func(c fiber.Ctx) error {
		ctx := c.Context()
		var request models.DeleteUserRequest

		if c.HasBody() {
			if err := c.Bind().Body(&request); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": placiderror.ErrBadRequest.Error()})
			}

			// Get in database user linked to email
			user, err := server.Queries.GetUser(ctx, pgtype.Text{
				Valid:  true,
				String: request.Email,
			})

			if err != nil {
				if errors.Is(err, pgx.ErrNoRows) {
					return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": placiderror.ErrUnauthorizedOperation.Error()})
				}
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": placiderror.ErrInternalServerError.Error()})
			}

			// Check if the password matches the hash
			err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash.String), []byte(request.Password))

			if err != nil {
				if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
					return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": placiderror.ErrUnauthorizedOperation.Error()})
				}

				server.Logger.Error(fmt.Sprintf("Email: %s - %s", request.Email, err.Error()))
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": placiderror.ErrInternalServerError.Error()})
			}

			err = server.Queries.DeleteUser(ctx, pgtype.Text{
				Valid:  true,
				String: request.Email,
			})

			if err != nil {
				server.Logger.Error(fmt.Sprintf("Email: %s - %s", request.Email, err.Error()))
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": placiderror.ErrInternalServerError.Error()})
			}

			return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Account deleted successfully"})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": placiderror.ErrBadRequest.Error()})

	}
}
