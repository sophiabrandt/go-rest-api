package http

import (
	"net/http"

	"github.com/sophiabrandt/go-rest-api/internal/data/book"
	"github.com/sophiabrandt/go-rest-api/internal/env"
	"github.com/sophiabrandt/go-rest-api/internal/transport/http/middleware"
)

// New creates a new router with all application routes.
func New(e *env.Env) http.Handler {

	bg := bookGroup{
		book: book.New(e.DB),
	}

	r := e.Router
	r.Handler(http.MethodGet, "/api/health", use(handler{e, health}, middleware.RequestLogger(e.Log)))
	r.Handler(http.MethodGet, "/api/books", use(handler{e, bg.getAllBooks}, middleware.RequestLogger(e.Log)))

	return r
}
