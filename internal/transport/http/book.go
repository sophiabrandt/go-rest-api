package http

import (
	"net/http"

	"github.com/sophiabrandt/go-rest-api/internal/data/book"
	"github.com/sophiabrandt/go-rest-api/internal/env"
)

type bookGroup struct {
	book *book.RepositoryDb
}

func (bg bookGroup) getAllBooks(e *env.Env, w http.ResponseWriter, r *http.Request) error {
	books, err := bg.book.Query()
	if err != nil {
		return err
	}
	return Respond(e, w, books, http.StatusOK)
}
