package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"placid-backend/internal/config"
	"placid-backend/internal/database"
	"placid-backend/internal/handlers"
	"placid-backend/internal/placid"

	"github.com/gofiber/fiber/v3"
	"github.com/redis/go-redis/v9"
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
			RedisConfig:    &placid.RedisConfig{},
			PostgresConfig: &placid.PostgresConfig{},
		},
	}

	if server.Cfg, err = config.FillEnv(); err != nil {
		server.Logger.Error(err.Error())
		os.Exit(1)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", server.Cfg.RedisConfig.RedisHost, server.Cfg.RedisConfig.RedisPort),
		Username: server.Cfg.RedisConfig.RedisUsername,
		Password: server.Cfg.RedisConfig.RedisPassword,
		DB:       0,
		Protocol: 2,
	})

	server.RedisClient = redisClient

	dbConn := config.ConnectPlacidDb(ctx, logger, server.Cfg.PostgresConfig)

	server.Queries = database.New(dbConn)
	server.App = app

	config.ConfigureCors(server.App, server.Cfg.AllowedCORSOrigins)

	// Register handlers
	handlers.RegisterRegisterHandler(server)
	handlers.RegisterLoginHandler(server)
	handlers.RegisterTracksHandler(server)
	handlers.RegisterUserHandler(server)
	handlers.RegisterDeleteAccountHandler(server)
	handlers.RegisterJoinNewsletterHandler(server)
	handlers.RegisterAdminUploadTrackHandler(server)
	handlers.RegisterAdminDeleteTracksHandler(server)

	if err = app.Listen(
		fmt.Sprintf(
			"%s:%s",
			server.Cfg.ApiConfig.ApiHost,
			server.Cfg.ApiConfig.ApiPort,
		)); err != nil {
		server.Logger.Error(err.Error())

		os.Exit(1)
	}
}
