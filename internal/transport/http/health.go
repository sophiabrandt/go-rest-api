package http

import (
	"net/http"

	"github.com/sophiabrandt/go-rest-api/internal/adapter/database"
	"github.com/sophiabrandt/go-rest-api/internal/env"
)

// health checks if the service is available and database is up.
func health(e *env.Env, w http.ResponseWriter, r *http.Request) error {
	status := "ok"
	statusCode := http.StatusOK
	if err := database.StatusCheck(e.DB); err != nil {
		status = "db not ready"
		statusCode = http.StatusInternalServerError
	}
	health := struct {
		Status string `json:"status`
	}{Status: status}

	return Respond(e, w, health, statusCode)
}
