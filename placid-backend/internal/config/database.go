package config

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"placid-backend/internal/placid"

	"github.com/jackc/pgx/v5"
)

func ConnectPlacidDb(ctx context.Context, logger *slog.Logger, dbCfg *placid.PostgresConfig) *pgx.Conn {
	connString := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		dbCfg.PostgresUsername,
		dbCfg.PostgresPassword,
		dbCfg.PostgresHost,
		dbCfg.PostgresPort,
		dbCfg.PostgresDb,
	)
	pgConn, err := pgx.Connect(ctx, connString)

	if err != nil {
		logger.Error(fmt.Sprintf("Failed to connect to database: %s\n", err.Error()))

		os.Exit(1)
	}

	// defer pgConn.Close(ctx)

	logger.Info("Database connection successful")

	return pgConn
}
