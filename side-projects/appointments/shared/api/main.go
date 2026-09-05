package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/luizandrends/appointments/shared/dependencies/postgres"
	"github.com/luizandrends/appointments/shared/dependencies/redis"
	"github.com/luizandrends/appointments/shared/routes"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Failed to execute server", "error", err)
		return
	}

	slog.Info("All Systems Offline")
}

func run() error {
	if err := postgres.RunDB(); err != nil {
		slog.Error("Failed to connect to postgres database", "error", err)
		return err
	}

	if err := redis.RunDB(); err != nil {
		slog.Error("Failed to connect to postgres database", "error", err)
		return err
	}

	routes := routes.NewHandler()

	s := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      routes,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
