package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"placid-backend/internal/auth"

	"github.com/jackc/pgx/v5"
)

type Server struct {
	dbConn *pgx.Conn
	mux    *http.ServeMux
	logger *slog.Logger
}

type User struct {
	Uuid  string
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (api *Server) run(addr string) error {
	err := http.ListenAndServe(addr, api.mux)

	if err != nil {
		return err
	}

	api.logger.Info(fmt.Sprintf("Server running at %s\n", addr))

	return nil
}

func (api *Server) registerHandlers() {
	auth.RegisterSignUpHandler(api.mux, api.logger)
	auth.RegisterSignInHandler(api.mux, api.logger)

}

func main() {
	server := Server{
		mux:    http.NewServeMux(),
		logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}

	server.registerHandlers()

	addr := fmt.Sprintf("%s:%d", "localhost", 8080)
	if err := server.run(addr); err != nil {
		server.logger.Error(err.Error())
		os.Exit(1)
	}
}
