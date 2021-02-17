package web

import (
	"net/http"
	"net/url"

	"github.com/pkg/errors"
	"github.com/sophiabrandt/go-rest-api/internal/data/book"
	"github.com/sophiabrandt/go-rest-api/internal/env"
)

type bookGroup struct {
	book *book.RepositoryDb
}

func (bg bookGroup) getRootBooksHandler(e *env.Env, w http.ResponseWriter, r *http.Request) error {
	// when searching the books api; example URI: /api/books?title="some+title"
	if len(params(r)) == 0 && len(r.URL.Query()) != 0 {
		bg.getBookSearch(e, w, r)
		return nil
	}
	bg.getAllBooks(e, w, r)
	return nil
}

func (bg bookGroup) getAllBooks(e *env.Env, w http.ResponseWriter, r *http.Request) error {
	books, err := bg.book.Query()
	if err != nil {
		return StatusError{err, http.StatusInternalServerError}
	}

	booksResp := books.ToDto()

	return respond(e, w, booksResp, http.StatusOK)
}

func (bg bookGroup) getBookSearch(e *env.Env, w http.ResponseWriter, r *http.Request) error {
	title, err := url.QueryUnescape(r.URL.Query().Get("title"))
	if err != nil {
		return StatusError{err, http.StatusBadRequest}
	}
	books, err := bg.book.QueryByTitle(title)
	if err != nil {
		switch errors.Cause(err) {
		case book.ErrNotFound:
			return StatusError{err, http.StatusNotFound}
		default:
			return errors.Wrapf(err, "Title : %s", title)
		}
	}

	booksResp := books.ToDto()

	return respond(e, w, booksResp, http.StatusOK)
}

func (bg bookGroup) getBookByID(e *env.Env, w http.ResponseWriter, r *http.Request) error {
	params := params(r)
	bk, err := bg.book.QueryByID(params["id"])
	if err != nil {
		switch errors.Cause(err) {
		case book.ErrInvalidID:
			return StatusError{err, http.StatusBadRequest}
		case book.ErrNotFound:
			return StatusError{err, http.StatusNotFound}
		default:
			return errors.Wrapf(err, "ID : %s", params["id"])
		}
	}
	bkResp := bk.ToDto()

	return respond(e, w, bkResp, http.StatusOK)
}

func (bg bookGroup) postBook(e *env.Env, w http.ResponseWriter, r *http.Request) error {
	var book book.NewBook

	if err := decode(r, &book); err != nil {
		return StatusError{err, http.StatusBadRequest}
	}

	if err := e.Validator.Struct(book); err != nil {
		resp := toErrResponse(err)
		return respond(e, w, resp, http.StatusUnprocessableEntity)
	}

	newBook, err := bg.book.Create(book)
	if err != nil {
		return StatusError{err, http.StatusInternalServerError}
	}

	newBookResp := newBook.ToDto()

	return respond(e, w, newBookResp, http.StatusCreated)
}

func (bg bookGroup) deleteBook(e *env.Env, w http.ResponseWriter, r *http.Request) error {
	params := params(r)

	if err := bg.book.Delete(params["id"]); err != nil {
		switch errors.Cause(err) {
		case book.ErrInvalidID:
			return StatusError{err, http.StatusBadRequest}
		default:
			return errors.Wrapf(err, "ID : %s", params["id"])
		}
	}

	return respond(e, w, nil, http.StatusNoContent)
}
