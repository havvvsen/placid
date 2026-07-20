package database

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5"
)

func ConnectDb(logger *slog.Logger) (*pgx.Conn, error) {
	pgConn, err := pgx.Connect(context.Background(), "postgresql://admin:default@localhost:5432/placid")
	if err != nil {
		return nil, err
	}
	logger.Info("Database connection successful")

	return pgConn, nil
}
