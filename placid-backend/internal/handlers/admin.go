package handlers

import (
	"errors"
	"fmt"
	"placid-backend/internal/database"
	"placid-backend/internal/models"
	"placid-backend/internal/placid"
	placiderror "placid-backend/internal/placid_error"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func RegisterAdminUploadTrackHandler(server *placid.Server) {
	server.App.Post(server.Cfg.Endpoints.AdminUploadTrack, func(c fiber.Ctx) error {
		ctx := c.Context()
		var request models.AdminUploadTrackRequest

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

				server.Logger.Info(fmt.Sprintf("Email: %s - %s", request.Email, err.Error()))
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": placiderror.ErrInternalServerError.Error()})
			}

			if user.IsAdmin.Valid && user.Token.Valid {
				if user.IsAdmin.Bool && user.Token.String == request.Token {
					err := server.Queries.AddTrack(
						ctx,
						database.AddTrackParams{
							Name:     request.Track.Name,
							Mood:     request.Track.Mood,
							AudioUrl: request.Track.AudioUrl,
							BgUrl:    request.Track.BgUrl,
						})

					if err != nil {
						server.Logger.Error(err.Error())
						return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": placiderror.ErrInternalServerError.Error()})
					}

					return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Tracks uploaded successfully"})
				}
			}

			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": placiderror.ErrUnauthorizedOperation.Error()})

		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": placiderror.ErrBadRequest.Error()})
	})
	server.Logger.Info(fmt.Sprintf("Registered %s", server.Cfg.Endpoints.AdminUploadTrack))
}

func RegisterAdminDeleteTracksHandler(server *placid.Server) {
	server.App.Delete(server.Cfg.Endpoints.AdminDeleteTrack, func(c fiber.Ctx) error {
		ctx := c.Context()
		var request models.AdminDeleteTrackRequest

		if c.HasBody() {
			if err := c.Bind().Body(&request); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": placiderror.ErrBadRequest.Error()})
			}

			user, err := server.Queries.GetUser(ctx, pgtype.Text{
				Valid:  true,
				String: request.Email,
			})

			if err != nil {
				server.Logger.Error(fmt.Sprintf("Email: %s - %s", request.Email, err.Error()))

				if errors.Is(err, pgx.ErrNoRows) {
					return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": placiderror.ErrUnauthorizedOperation.Error()})
				}

				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": placiderror.ErrInternalServerError.Error()})
			}

			if user.IsAdmin.Valid && user.Token.Valid {
				if user.IsAdmin.Bool && user.Token.String == request.Token {
					err = server.Queries.DeleteTrack(ctx, request.Track.ID)

					if err != nil {
						server.Logger.Error(fmt.Sprintf("Email: %s - %s", request.Email, err.Error()))
						return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": placiderror.ErrInternalServerError.Error()})
					}

					return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Tracks deleted successfully"})
				}
			}

			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": placiderror.ErrUnauthorizedOperation.Error()})

		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": placiderror.ErrBadRequest.Error()})
	})
	server.Logger.Info(fmt.Sprintf("Registered %s", server.Cfg.Endpoints.AdminDeleteTrack))
}
