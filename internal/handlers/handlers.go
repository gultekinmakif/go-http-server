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
func writeJSON(w http.ResponseWriter, status int, v any) {
	writeHeaders(w, status)
	json.NewEncoder(w).Encode(v)
}

func Root(w http.ResponseWriter, r *http.Request) {
	writeHeaders(w, http.StatusOK)
}

func Health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, healthResponse{Status: "ok"})
}
