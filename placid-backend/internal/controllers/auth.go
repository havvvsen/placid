package controllers

import (
	"errors"
	"fmt"
	"placid-backend/internal/database"
	"placid-backend/internal/models"
	"placid-backend/internal/placid"
	placiderror "placid-backend/internal/placid_error"
	"placid-backend/internal/utils"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
)

func ControllerRegister(server *placid.Server) fiber.Handler {
	return func(c fiber.Ctx) error {
		ctx := c.Context()
		var request models.RegisterRequest
		var pgErr *pgconn.PgError

		if c.HasBody() {
			if err := c.Bind().Body(&request); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": placiderror.ErrBadRequest.Error()})
			}

			// TODO: Do proper input sanitization
			err := utils.SanitizeAuthRequest(&request)

			if err != nil {
				if errors.Is(err, placiderror.ErrInsecurePassword) || errors.Is(err, placiderror.ErrInvalidEmail) {
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
				}

				server.Logger.Error(fmt.Sprintf("Failed to sanitize auth request: %s", err.Error()))
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": placiderror.ErrBadRequest.Error()})
			}

			passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

			if err != nil {
				server.Logger.Error(fmt.Sprintf("Failed to hash password: %s", err.Error()))
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Registration failed"})
			}

			err = server.Queries.AddUser(ctx, database.AddUserParams{
				Email: pgtype.Text{
					Valid:  true,
					String: request.Email,
				},
				PasswordHash: pgtype.Text{
					Valid:  true,
					String: string(passwordHash),
				},
			})

			if err != nil {
				// Duplicate key violation. User already exists
				if errors.As(err, &pgErr) && pgErr.Code == "23505" {
					return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": fmt.Sprintf("Email %s is already in use", request.Email)})
				}

				server.Logger.Error(fmt.Sprintf("Failed to sign up user: %s", err.Error()))

				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Registration failed"})

			}

			return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Registration successful"})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Request has no body"})
	}
}

func ControllerLogin(server *placid.Server) fiber.Handler {
	return func(c fiber.Ctx) error {
		ctx := c.Context()
		var request models.LoginRequest

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
					// User does not exist in database
					return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": placiderror.ErrInvalidCredentials.Error()})
				}

				server.Logger.Error(fmt.Sprintf("%s - %s", request.Email, err.Error()))
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": placiderror.ErrInternalServerError.Error()})

			}

			err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash.String), []byte(request.Password))

			if err != nil {
				if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
					return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": placiderror.ErrInvalidCredentials.Error()})
				}

				server.Logger.Error(fmt.Sprintf("%s - %s", request.Email, err.Error()))
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": placiderror.ErrInternalServerError.Error()})
			}

			token, err := utils.GenerateJWTToken(server.Cfg.Secrets.JwtSecretKey, user.Email.String, user.IsAdmin.Bool)

			if err != nil {
				server.Logger.Error(fmt.Sprintf("Failed to sign jwt token: %s", err.Error()))
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": placiderror.ErrInternalServerError.Error()})
			}

			return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Authentication successful", "token": token})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": placiderror.ErrBadRequest.Error()})
	}

}
