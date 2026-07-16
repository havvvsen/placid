package database

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5"
)

func ConnectDb(trials int, logger *slog.Logger) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), "")

	if err != nil {
		return nil, err
	}

	logger.Info("Database connection successful")

	return conn, nil
}
