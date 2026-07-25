package handlers

import (
	"fmt"
	"placid-backend/internal/placid"
	placiderror "placid-backend/internal/placid_error"

	"github.com/gofiber/fiber/v3"
)

func RegisterTracksHandler(server *placid.Server) {
	server.App.Get(server.Cfg.Endpoints.Tracks, func(c fiber.Ctx) error {
		ctx := c.Context()

		tracks, err := server.Queries.GetTracks(ctx)

		if err != nil {
			server.Logger.Error(err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": placiderror.ErrInternalServerError.Error()})
		}

		return c.Status(fiber.StatusOK).JSON(tracks)
	})

	server.Logger.Info(fmt.Sprintf("Registered %s", server.Cfg.Endpoints.Tracks))
}
