package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"placid-backend/internal/auth"

	"github.com/jackc/pgx/v5"
)

type Api struct {
	dbConn *pgx.Conn
	logger *slog.Logger
	mux    *http.ServeMux
	host   string
	port   uint16
}

func (api *Api) run() error {
	addr := fmt.Sprintf("%s:%d", api.host, api.port)
	err := http.ListenAndServe(addr, api.mux)

	if err != nil {
		return err
	}

	api.logger.Info(fmt.Sprintf("Server running at %s\n", addr))

	return nil
}

func (api *Api) registerHandlers() {
	auth.RegisterSignUpHandler(api.mux, api.logger)
	auth.RegisterSignInHandler(api.mux, api.logger)

}

func main() {
	api := Api{
		mux:    http.NewServeMux(),
		logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
		host:   "localhost",
		port:   8080,
	}

	api.registerHandlers()

	if err := api.run(); err != nil {
		api.logger.Error(err.Error())
		os.Exit(1)
	}
}
