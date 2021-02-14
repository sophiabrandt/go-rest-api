package http

import (
	"net/http"

	"github.com/justinas/alice"
	"github.com/sophiabrandt/go-rest-api/internal/env"
	"github.com/sophiabrandt/go-rest-api/internal/transport/http/middleware"
)

// New creates a new router with all application routes.
func New(e *env.Env) http.Handler {
	standardMiddleware := alice.New(middleware.RequestLogger(e.Log))

	r := e.Router
	r.Handler(http.MethodGet, "/api/health", Handler{e, HandleHealth})

	return standardMiddleware.Then(r)
}
