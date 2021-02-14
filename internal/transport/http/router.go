package http

import (
	"net/http"

	"github.com/sophiabrandt/go-rest-api/internal/env"
)

// New creates a new router with all application routes.
func New(e *env.Env) http.Handler {
	r := e.Router
	r.Handler(http.MethodGet, "/api/health", Handler{e, HandleHealth})

	return r
}
