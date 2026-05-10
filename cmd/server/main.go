package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/gultekinmakif/go-http-server/internal/config"
	"github.com/gultekinmakif/go-http-server/internal/logger"
)

func roothandler(res http.ResponseWriter, req *http.Request)   {}
func healthHandler(res http.ResponseWriter, req *http.Request) {}

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
	mux.HandleFunc("/", roothandler)
	mux.HandleFunc("/health", healthHandler)

	if err := http.ListenAndServe(":"+cfg.Port, mux); err != nil {
		slog.Error("server error", "error", err)
	}
}
