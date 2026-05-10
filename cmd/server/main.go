package main

import (
	"context"
	"log"
	"log/slog"
	"os/signal"
	"syscall"

	"github.com/gultekinmakif/go-http-server/internal/config"
	"github.com/gultekinmakif/go-http-server/internal/logger"
	"github.com/gultekinmakif/go-http-server/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	lg, err := logger.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	slog.SetDefault(lg)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := server.New(cfg)
	if err := srv.Start(ctx); err != nil {
		slog.Error("server error", "error", err)
	}
}
