package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	basePath := "/home/hansen/Projects/placid/tracks-server/assets"
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	server := http.FileServer(http.Dir(basePath))

	if err := http.ListenAndServe(":3031", server); err != nil {
		logger.Info("Server serving at :3031")
		fmt.Println("Here")
		log.Fatal(err.Error())
	}
}
