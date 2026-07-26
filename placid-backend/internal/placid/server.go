package placid

import (
	"log/slog"
	"placid-backend/internal/database"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5"
)

type Server struct {
	Cfg     *Config
	App     *fiber.App
	DbConn  *pgx.Conn
	Queries *database.Queries
	Logger  *slog.Logger
}
