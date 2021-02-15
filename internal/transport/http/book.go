package http

import (
	"encoding/json"
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
		return StatusError{404, err}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
	return nil
}
