package controllers

import (
	"placid-backend/internal/placid"
	placiderror "placid-backend/internal/placid_error"

	"github.com/gofiber/fiber/v3"
)

func ControllerTracks(server *placid.Server) fiber.Handler {
	return func(c fiber.Ctx) error {
		ctx := c.Context()

		tracks, err := server.Queries.GetTracks(ctx)

		if err != nil {
			server.Logger.Error(err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": placiderror.ErrInternalServerError.Error()})
		}

		return c.Status(fiber.StatusOK).JSON(tracks)
	}

}
