package api

import (
	"encoding/json"
	"net/http"
)

// Healthy healt status check
func Healthy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Healthy")
}
