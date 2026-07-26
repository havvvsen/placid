package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"placid-backend/internal/config"
	"placid-backend/internal/controllers"
	"placid-backend/internal/database"
	"placid-backend/internal/initializers"
	"placid-backend/internal/middlewares"
	"placid-backend/internal/placid"

	"github.com/gofiber/fiber/v3"
)

func main() {
	ctx := context.Background()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	var err error

	app := fiber.New()

	server := &placid.Server{
		Logger: logger,
		Cfg: &placid.Config{
			ApiConfig:      &placid.ApiConfig{},
			PostgresConfig: &placid.PostgresConfig{},
		},
	}

	if server.Cfg, err = initializers.FillEnv(); err != nil {
		server.Logger.Error(err.Error())
		os.Exit(1)
	}

	dbConn := initializers.ConnectDatabase(ctx, logger, server.Cfg.PostgresConfig)

	server.Queries = database.New(dbConn)
	server.App = app

	config.ConfigureCors(server.App, server.Cfg.AllowedCORSOrigins)

	// Register controllers
	// Register
	server.App.Post(server.Cfg.Endpoints.Register, controllers.ControllerRegister(server))
	// Login
	server.App.Post(server.Cfg.Endpoints.Login, controllers.ControllerLogin(server))
	// Users
	server.App.Post(
		server.Cfg.Endpoints.User,
		middlewares.AuthMiddleware(
			server.Logger,
			server.Cfg.Secrets.JwtSecretKey,
		),
		controllers.ControllerUser(server),
	)
	// Delete Account
	server.App.Delete(
		server.Cfg.Endpoints.DeleteAccount,
		middlewares.AuthMiddleware(
			server.Logger,
			server.Cfg.Secrets.JwtSecretKey,
		),
		controllers.ControllerDeleteAccount(server),
	)
	// Tracks
	server.App.Get(server.Cfg.Endpoints.Tracks,
		middlewares.AuthMiddleware(
			server.Logger,
			server.Cfg.Secrets.JwtSecretKey,
		),
		controllers.ControllerTracks(server),
	)
	// Subscribe Newsletter
	server.App.Post(server.Cfg.Endpoints.SubscribeNewsletter, controllers.ControllerSubscribeNewsletter(server))
	// Unsubscribe Newsletter
	server.App.Delete(server.Cfg.Endpoints.UnsubscribeNewsletter, controllers.ControllerUnsubscribeNewsletter(server))
	// Admin Upload Track
	server.App.Post(server.Cfg.Endpoints.AdminUploadTrack,
		middlewares.AuthMiddleware(
			server.Logger,
			server.Cfg.Secrets.JwtSecretKey,
		),
		controllers.ControllerAdminUploadTrack(server),
	)
	// Admin Delete Track
	server.App.Delete(
		server.Cfg.Endpoints.AdminDeleteTrack,
		middlewares.AuthMiddleware(
			server.Logger,
			server.Cfg.Secrets.JwtSecretKey,
		),

		controllers.ControllerAdminDeleteTrack(server))

	// if err = app.Listen(
	// 	fmt.Sprintf(
	// 		"%s:%s",
	// 		server.Cfg.ApiConfig.ApiHost,
	// 		server.Cfg.ApiConfig.ApiPort,
	// 	)); err != nil {
	// 	server.Logger.Error(err.Error())
	//
	// 	os.Exit(1)
	// }
	if err = app.Listen(server.Cfg.ApiConfig.ApiPort); err != nil {
		server.Logger.Error(err.Error())

		os.Exit(1)
	}
}
