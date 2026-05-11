package main

import (
	"context"
	"log"
	"log/slog"
	"os/signal"
	"syscall"

	"github.com/gultekinmakif/go-http-server/internal/config"
	"github.com/gultekinmakif/go-http-server/internal/db/postgres"
	"github.com/gultekinmakif/go-http-server/internal/logger"
	"github.com/gultekinmakif/go-http-server/internal/server"
)

func main() {
	defer postgres.Close()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	lg, err := logger.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	slog.SetDefault(lg)

	if err := postgres.New(cfg.DatabaseURL); err != nil {
		log.Fatal(err)
	}

	if err := postgres.Migrate(); err != nil {
		log.Fatal(err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := server.New(cfg)
	if err := srv.Start(ctx); err != nil {
		slog.Error("server error", "error", err)
	}
}
