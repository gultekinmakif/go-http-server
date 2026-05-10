package handlers

import (
	"encoding/json"
	"net/http"
)

type healthResponse struct {
	Status string `json:"status"`
}

func writeHeaders(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
}

func Root(w http.ResponseWriter, r *http.Request) {
	writeHeaders(w, http.StatusOK)
}

func Health(w http.ResponseWriter, r *http.Request) {
	writeHeaders(w, http.StatusOK)
	json.NewEncoder(w).Encode(healthResponse{Status: "ok"})
}
