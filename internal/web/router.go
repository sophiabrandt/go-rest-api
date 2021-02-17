package web

import (
	"net/http"

	"github.com/sophiabrandt/go-rest-api/internal/data/book"
	"github.com/sophiabrandt/go-rest-api/internal/env"
	"github.com/sophiabrandt/go-rest-api/internal/web/middleware"
)

// NewRouter creates a new router with all application routes.
func NewRouter(e *env.Env) http.Handler {
	bg := bookGroup{
		book: book.New(e.Db),
	}

	r := e.Router
	r.Handler(http.MethodGet, "/api/health", use(handler{e, health}, middleware.RequestLogger(e.Log)))
	r.Handler(http.MethodGet, "/api/books", use(handler{e, bg.getRootBooksHandler}, middleware.RequestLogger(e.Log)))
	r.Handler(http.MethodGet, "/api/books/:id", use(handler{e, bg.getBookByID}, middleware.RequestLogger(e.Log)))
	r.Handler(http.MethodPost, "/api/books", use(handler{e, bg.postBook}, middleware.RequestLogger(e.Log)))

	return r
}
