package main

import (
	"aantonioprado/rs-go-api-movie-search/api"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Failed to execute code", "error", err)
		os.Exit(1)
	}

	slog.Info("Code executed successfully")
}

func run() error {
	apiKey := os.Getenv("APIKEY_OMDB")
	handler := api.NewHandler(apiKey)

	s := http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Minute,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
