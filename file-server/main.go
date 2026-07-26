package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	basePath := "/home/ubuntu/placid/file-server/assets"
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	port, isSet := os.LookupEnv("FILE_SERVER_PORT")

	if !isSet {
		logger.Error("$FILE_SERVER_PORT is required")
		os.Exit(1)
	}

	addr := fmt.Sprintf(":%s", port)

	logger.Info(fmt.Sprintf("Serving assets at %s", addr))

	server := http.FileServer(http.Dir(basePath))

	if err := http.ListenAndServe(addr, server); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
