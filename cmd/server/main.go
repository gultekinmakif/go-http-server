package main

import (
	"log"
	"net/http"
	"os"
)

func roothandler(res http.ResponseWriter, req *http.Request)   {}
func healthHandler(res http.ResponseWriter, req *http.Request) {}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("server listening on %s", port)

	mux := http.NewServeMux()
	mux.HandleFunc("/", roothandler)
	mux.HandleFunc("/health", healthHandler)

	http.ListenAndServe(":"+port, mux)
}
