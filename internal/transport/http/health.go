package http

import (
	"encoding/json"
	"net/http"

	"github.com/sophiabrandt/go-rest-api/internal/env"
)

func HandleHealth(e *env.Env, w http.ResponseWriter, r *http.Request) error {
	status := "ok"
	statusCode := http.StatusOK
	health := struct {
		Status string `json:"status`
	}{Status: status}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(health)
	return nil
}
