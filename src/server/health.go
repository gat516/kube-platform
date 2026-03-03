package server

import (
	"encoding/json"
	"net/http"
	"time"
)

// healthResponse is the JSON body returned by the health endpoint.
type healthResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Version   string `json:"version"`
}

// handleHealth responds with the service liveness status.
// It is intentionally lightweight — no external dependency checks — so that
// Kubernetes liveness probes get a fast, reliable response.
func handleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(healthResponse{
		Status:    "ok",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Version:   "0.1.0",
	})
}
