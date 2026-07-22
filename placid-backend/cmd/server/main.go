package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"
	"placid-backend/internal/database"
	"placid-backend/internal/models"
	placiderror "placid-backend/internal/placid_error"
	"placid-backend/pkg/utils"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

type Server struct {
	app       *fiber.App
	pgConn    *pgx.Conn
	logger    *slog.Logger
	rdbClient *redis.Client
}

func main() {
	app := fiber.New()

	app.Use(
		cors.New(
			cors.Config{
				AllowOrigins: []string{"http://localhost:4200"},
				AllowMethods: []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
				AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
			},
		),
	)

	rdbClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		Protocol: 2,
	})

	server := Server{
		app:       app,
		logger:    slog.New(slog.NewTextHandler(os.Stdout, nil)),
		rdbClient: rdbClient,
	}

	pgConn, err := database.ConnectDb(server.logger)

	if err != nil {
		server.logger.Error(fmt.Sprintf("Failed to connect to database: %s\n", err.Error()))
	}

	server.pgConn = pgConn

	app.Post("/api/v1/auth/register", func(c fiber.Ctx) error {
		var user models.User
		var pgErr *pgconn.PgError

		if c.HasBody() {
			if err := c.Bind().Body(&user); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": placiderror.ErrBadRequest.Error()})
			}

			err := utils.SanitizeAuthRequest(&user)

			if err != nil {
				if errors.Is(err, placiderror.ErrInsecurePassword) || errors.Is(err, placiderror.ErrInvalidEmail) {
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
				}

				server.logger.Error(fmt.Sprintf("Failed to sanitize auth request: %s\n", err.Error()))
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": placiderror.ErrBadRequest.Error()})
			}

			passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

			if err != nil {
				server.logger.Error(fmt.Sprintf("Failed to hash password: %s\n", err.Error()))
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Registration failed"})
			}

			// Sign up user
			_, err = server.pgConn.Exec(context.Background(), "INSERT INTO users ( email, password ) VALUES ( $1, $2 );", user.Email, passwordHash)

			if err != nil {
				// Duplicate key violation. User already exists
				if errors.As(err, &pgErr) && pgErr.Code == "23505" {
					return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": fmt.Sprintf("Email %s is already in use\n", user.Email)})
				}

				server.logger.Error(fmt.Sprintf("Failed to sign up user: %s\n", err.Error()))

				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Registration failed"})

			}

			return c.JSON(fiber.Map{"message": "Registration successful"})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Request has no body"})
	})

	app.Post("/api/v1/auth/login", func(c fiber.Ctx) error {
		var user models.User
		var inDatabaseHashedPassword []byte

		if c.HasBody() {
			if err := c.Bind().Body(&user); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": placiderror.ErrBadRequest.Error()})
			}

			row := pgConn.QueryRow(context.Background(), "SELECT password FROM users WHERE email=$1;", user.Email)

			if err := row.Scan(&inDatabaseHashedPassword); err != nil {
				if errors.Is(err, pgx.ErrNoRows) {
					// User does not exist in database
					return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": placiderror.ErrInvalidCredentials.Error()})
				}

				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": placiderror.ErrInternalServerError.Error()})
			}

			err = bcrypt.CompareHashAndPassword(inDatabaseHashedPassword, []byte(user.Password))

			if err != nil {
				if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
					return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": placiderror.ErrInvalidCredentials.Error()})
				}

				server.logger.Error(fmt.Sprintf("Compare hash and pass: %s\n", err.Error()))
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": placiderror.ErrInternalServerError.Error()})
			}

			return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Authentication successful", "token": "jwt-token-8762y8a82n0824bsbfy82"})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": placiderror.ErrBadRequest.Error()})
	})

	app.Delete("/api/v1/delete-account", func(c fiber.Ctx) error {
		var user models.User
		var inDatabaseHashedPassword []byte

		if c.HasBody() {
			if err := c.Bind().Body(&user); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": placiderror.ErrBadRequest.Error()})
			}

			rows, err := pgConn.Query(context.Background(), "SELECT password FROM users WHERE email==$1;", user.Email)

			if errors.Is(err, pgx.ErrNoRows) {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": placiderror.ErrInvalidCredentials.Error()})
			}

			if err := rows.Scan(&inDatabaseHashedPassword); err != nil {
				server.logger.Error(fmt.Sprintf("Failed to scan password hash: %s\n", err.Error()))
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": placiderror.ErrInternalServerError.Error()})
			}

			err = bcrypt.CompareHashAndPassword(inDatabaseHashedPassword, []byte(user.Password))

			if err != nil {
				if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
					return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": placiderror.ErrInvalidCredentials.Error()})
				}

				server.logger.Error(fmt.Sprintf("Failed to compare hash and password: %s\n", err.Error()))
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": placiderror.ErrInternalServerError.Error()})
			}

			_, err = pgConn.Exec(context.Background(), "DELETE FROM users WHERE email = $1;", user.Email)

			if err != nil {
				server.logger.Error(fmt.Sprintf("Failed to delete account: %s\n", err.Error()))
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete account"})

			}

			c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Account deleted successfully"})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": placiderror.ErrBadRequest.Error()})

	})

	app.Post("/api/v1/admin/upload", func(c fiber.Ctx) error {

		return c.SendString("H")
	})

	app.Get("/api/v1/soundscapes", func(c fiber.Ctx) error {
		// rows, err := pgConn.Query(context.Background(), "SELECT * FROM soundscapes;")

		// if err != nil {
		// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": placiderror.ErrInternalServerError})

		// }

		// soundScapes, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Soundscape])

		// if err != nil {
		// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": placiderror.ErrInternalServerError})
		// }

		soundScapes := `
		{
		"focus": [
			{
			"name": "Rain noise"
			"url": "https://placid-soundscapes-s3-bucket.s3.us-east-1.amazonaws.com/relax/4b.mp3"
			},
			{
			"name": "Rain noise"
			"url": "https://placid-soundscapes-s3-bucket.s3.us-east-1.amazonaws.com/relax/4b.mp3"
			},
			{
			"name": "Rain noise"
			"url": "https://placid-soundscapes-s3-bucket.s3.us-east-1.amazonaws.com/relax/4b.mp3"
			},

		],
		"sleep": [
		{
		"name": "Ocean Waves"
		"url": "https://placid-soundscapes-s3-bucket.s3.us-east-1.amazonaws.com/relax/4b.mp3"
		},
		{
		"name": "Ocean Waves"
		"url": "https://placid-soundscapes-s3-bucket.s3.us-east-1.amazonaws.com/relax/4b.mp3"
		},
		{
		"name": "Ocean Waves"
		"url": "https://placid-soundscapes-s3-bucket.s3.us-east-1.amazonaws.com/relax/4b.mp3"
		},

		],
		"relax": [
		{
		"name": "Ocean Waves"
		"url": "https://placid-soundscapes-s3-bucket.s3.us-east-1.amazonaws.com/relax/4b.mp3"
		},
		{
		"name": "Ocean Waves"
		"url": "https://placid-soundscapes-s3-bucket.s3.us-east-1.amazonaws.com/relax/4b.mp3"
		},

		]
		}
		`

		return c.Status(fiber.StatusOK).JSON(strings.Trim(soundScapes, " "))
	})

	app.Delete("/api/v1/admin/delete-sound", func(c fiber.Ctx) error {
		return c.SendString("H")
	})

	app.Post("/api/v1/join-newsletter", func(c fiber.Ctx) error {
		var request models.BasicEmailModel
		var disposable models.BasicEmailModel
		if c.HasBody() {
			if err := c.Bind().Body(&request); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": placiderror.ErrBadRequest.Error()})
			}

			row := pgConn.QueryRow(context.Background(), "SELECT email FROM newsletter WHERE email = $1;", strings.Trim(request.Email, " "))
			err := row.Scan(&disposable.Email)

			if err != nil {
				if errors.Is(err, pgx.ErrNoRows) {
					query := "INSERT INTO newsletter (email) VALUES ( $1 );"
					_, err = pgConn.Exec(context.Background(), query, request.Email)

					if err != nil {
						server.logger.Error(fmt.Sprintf("Failed to add email to newsletter: %s\n", err.Error()))
						return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to add email to newsletter"})
					}

					return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Successful added to newsletter"})
				}

				server.logger.Error(fmt.Sprintf("Failed to add email to newsletter: %s\n", err.Error()))
				return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message": placiderror.ErrInternalServerError.Error()})
			}

			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": placiderror.ErrEmailExistsNewsletter.Error()})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Bad Request"})
	})

	log.Fatal(app.Listen(":3000"))

}
