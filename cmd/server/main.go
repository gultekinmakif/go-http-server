package main

import (
	"log"
	"net/http"

	"github.com/gultekinmakif/go-http-server/internal/config"
)

func roothandler(res http.ResponseWriter, req *http.Request)   {}
func healthHandler(res http.ResponseWriter, req *http.Request) {}

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("server listening on %s (env=%s)", cfg.Port, cfg.Env)

	mux := http.NewServeMux()
	mux.HandleFunc("/", roothandler)
	mux.HandleFunc("/health", healthHandler)

	if err := http.ListenAndServe(":"+cfg.Port, mux); err != nil {
		log.Fatal(err)
	}
}
