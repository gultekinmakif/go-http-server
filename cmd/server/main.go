package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/gultekinmakif/go-http-server/internal/config"
	"github.com/gultekinmakif/go-http-server/internal/handlers"
	"github.com/gultekinmakif/go-http-server/internal/logger"
	"github.com/gultekinmakif/go-http-server/internal/middleware"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	_logger, err := logger.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	slog.SetDefault(_logger)

	slog.Info("server listening", "port", cfg.Port, "env", cfg.Env)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Root)
	mux.HandleFunc("/health", handlers.Health)

	router := middleware.Recoverer(middleware.RequestID(middleware.Logger(mux)))

	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		slog.Error("server error", "error", err)
	}
}
