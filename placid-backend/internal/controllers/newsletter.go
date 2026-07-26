package controllers

import (
	"errors"
	"fmt"
	"placid-backend/internal/models"
	"placid-backend/internal/placid"
	placiderror "placid-backend/internal/placid_error"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5"
)

func ControllerSubscribeNewsletter(server *placid.Server) fiber.Handler {
	return func(c fiber.Ctx) error {
		ctx := c.Context()
		var request models.SubscribeToNewsletterRequest

		if c.HasBody() {
			if err := c.Bind().Body(&request); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": placiderror.ErrBadRequest.Error()})
			}

			_, err := server.Queries.GetNewsletterSubscriber(ctx, request.Email)

			if err != nil {
				if errors.Is(err, pgx.ErrNoRows) {
					err := server.Queries.AddNewsletterSubsciber(ctx, request.Email)

					if err != nil {
						server.Logger.Error(fmt.Sprintf("Failed to add email to newsletter: %s", err.Error()))
						return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": placiderror.ErrInternalServerError.Error()})
					}

					return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Successful added to newsletter"})
				}

				server.Logger.Error(fmt.Sprintf("Failed to add email to newsletter: %s", err.Error()))
				return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message": placiderror.ErrInternalServerError.Error()})
			}

			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": placiderror.ErrEmailExistsNewsletter.Error()})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": placiderror.ErrBadRequest.Error()})
	}
}

func ControllerUnsubscribeNewsletter(server *placid.Server) fiber.Handler {
	return func(c fiber.Ctx) error {
		ctx := c.Context()

		var request models.UnsubscribeFromNewsletterRequest

		if c.HasBody() {
			if err := c.Bind().Body(&request); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": placiderror.ErrBadRequest.Error()})
			}

			_, err := server.Queries.GetNewsletterSubscriber(ctx, request.Email)

			if err != nil {
				if errors.Is(err, pgx.ErrNoRows) {
					return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "You are currently not subscribed to our newsletter. Please consider subscribing"})
				}

				server.Logger.Error(fmt.Sprintf("Failed to remove email from newsletter: %s", err.Error()))
				return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message": placiderror.ErrInternalServerError.Error()})
			}

			err = server.Queries.DeleteNewsletterSubscriber(ctx, request.Email)

			if err != nil {
				server.Logger.Error(fmt.Sprintf("Failed to remove email from newsletter: %s", err.Error()))
				return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message": placiderror.ErrInternalServerError.Error()})
			}

			return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": "You have successfully unsubscribed from our newsletter"})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": placiderror.ErrBadRequest.Error()})

	}
}
