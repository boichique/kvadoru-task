package main

import (
	"context"
	"net/http"
	"os"

	"github.com/boichique/kvadoru_task/internal/config"
	"github.com/boichique/kvadoru_task/internal/server"
	"golang.org/x/exp/slog"
)

func main() {
	cfg, err := config.NewConfig() // Initialize a new configuration
	failOnError(err, "parse config")

	srv, err := server.NewServer(context.Background(), cfg) // Initialize a new server
	failOnError(err, "create server")

	// Start the server, if there is an error and it's not ErrServerClosed, log it and exit
	if err := srv.Start(); err != nil && err != http.ErrServerClosed {
		slog.Error(
			"server stopped",
			"error", err,
		)
		os.Exit(1)
	}
}

// Function to handle errors: if there is an error, log it and exit
func failOnError(err error, msg string) {
	if err != nil {
		slog.Error(
			msg,
			"error", err,
		)
		os.Exit(1)
	}
}
