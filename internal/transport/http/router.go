package http

import (
	"net/http"

	"github.com/justinas/alice"
	"github.com/sophiabrandt/go-rest-api/internal/data/book"
	"github.com/sophiabrandt/go-rest-api/internal/env"
	"github.com/sophiabrandt/go-rest-api/internal/transport/http/middleware"
)

// New creates a new router with all application routes.
func New(e *env.Env) http.Handler {
	standardMiddleware := alice.New(middleware.RequestLogger(e.Log))

	bg := bookGroup{
		book: book.New(e.DB),
	}

	r := e.Router
	r.Handler(http.MethodGet, "/api/health", Handler{e, Health})

	r.Handler(http.MethodGet, "/api/books", Handler{e, bg.getAllBooks})

	return standardMiddleware.Then(r)
}
