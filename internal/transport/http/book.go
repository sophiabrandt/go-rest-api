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
		return StatusError{err, http.StatusInternalServerError}
	}
	return Respond(e, w, books, http.StatusOK)
}

func (bg bookGroup) PostBook(e *env.Env, w http.ResponseWriter, r *http.Request) error {
	var book book.NewBook
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		return StatusError{err, http.StatusBadRequest}
	}
	newBook, err := bg.book.Create(book)
	if err != nil {
		return StatusError{err, http.StatusInternalServerError}
	}
	return Respond(e, w, newBook, http.StatusCreated)
}
