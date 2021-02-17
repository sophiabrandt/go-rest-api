package book

import (
	"database/sql"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

var (
	// ErrNotFound is used when a specific Product is requested but does not exist.
	ErrNotFound = errors.New("not found")

	// ErrInvalidID occurs when an ID is not in a valid form.
	ErrInvalidID = errors.New("ID is not in its proper form")
)

// BookRepositoryDb defines the repository for the book service.
type RepositoryDb struct {
	Db *sqlx.DB
}

// Repo is the interface for the book repository.
type Repo interface {
	Query() (Infos, error)
	QueryByID(bookID string) (Info, error)
	QueryByTitle(bookTitle string) (Info, error)
	Create(book NewBook) (Info, error)
	Delete(bookID string) error
}

// New returns a pointer to a book repo.
func New(db *sqlx.DB) *RepositoryDb {
	return &RepositoryDb{Db: db}
}

// Query retrieves all books from the database.
func (r *RepositoryDb) Query() (Infos, error) {
	const q = `
	SELECT
		b.book_id, b.title, b.published_date, b.image_url, b.description,
		a.name AS author_name
	FROM books as b
	LEFT JOIN
		authors AS a ON b.author_id = a.author_id
	ORDER BY b.book_id
	`
	var books Infos
	if err := r.Db.Select(&books, q); err != nil {
		return books, errors.Wrap(err, "selecting books")
	}
	return books, nil
}

// QuerybyID retrieves a book by ID from the database.
func (r *RepositoryDb) QueryByID(bookID string) (Info, error) {
	if _, err := uuid.Parse(bookID); err != nil {
		return Info{}, ErrInvalidID
	}

	const q = `
	SELECT
		b.book_id, b.title, b.published_date, b.image_url, b.description,
		a.name AS author_name
	FROM books as b
	LEFT JOIN
		authors AS a ON b.author_id = a.author_id
	WHERE
		b.book_id = $1
	`

	var book Info
	if err := r.Db.Get(&book, q, bookID); err != nil {
		if err == sql.ErrNoRows {
			return book, ErrNotFound
		}
		return book, errors.Wrapf(err, "selecting book with ID %s", bookID)
	}
	return book, nil
}

// QuerybyTitle retrieves a book by quering the title from the database.
func (r *RepositoryDb) QueryByTitle(bookTitle string) (Infos, error) {
	const q = `
	SELECT
		b.book_id, b.title, b.published_date, b.image_url, b.description,
		a.name AS author_name
	FROM books as b
	LEFT JOIN
		authors AS a ON b.author_id = a.author_id
	WHERE
		b.title LIKE '%' || $1 || '%'
	`
	var books []Info
	if err := r.Db.Select(&books, q, bookTitle); err != nil {
		if err == sql.ErrNoRows {
			return books, ErrNotFound
		}
		return books, errors.Wrap(err, "selecting book by Title")
	}
	return books, nil
}

// Create adds a new book to the database. It returns the created book with
// fields like ID and Author_ID populated.
func (r *RepositoryDb) Create(book NewBook) (Info, error) {
	// find the id for the author
	const a_id = `
	SELECT a.author_id
	FROM
		authors AS a
	WHERE a.name = $1
	`
	var author_id string
	if err := r.Db.Get(&author_id, a_id, book.AuthorName); err != nil {
		// author does not exist yet, create new author
		author_id = uuid.New().String()

		const a = `
		INSERT INTO authors
			(author_id, name)
		VALUES
		($1, $2)
		`

		if _, err := r.Db.Exec(a, author_id, book.AuthorName); err != nil {
			return Info{}, errors.Wrap(err, "selecting author for book")
		}
	}

	pubDate, err := time.Parse("2006-01-02", book.PublishedDate)
	if err != nil {
		return Info{}, errors.Wrap(err, "parsing date format for published_date")
	}
	dateString := pubDate.Format("2006-01-02")

	// create new book model for the database
	bk := Info{
		ID:            uuid.New().String(),
		AuthorID:      author_id,
		AuthorName:    strings.ToLower(book.AuthorName),
		Title:         strings.ToLower(book.Title),
		PublishedDate: dateString,
		ImageUrl:      book.ImageUrl,
		Description:   book.Description,
	}

	const q = `
	INSERT INTO books
		(book_id, title, author_id, published_date, image_url, description)
	VALUES
		($1, $2, $3, $4, $5, $6)`

	if _, err := r.Db.Exec(q, bk.ID, bk.Title, bk.AuthorID, bk.PublishedDate, bk.ImageUrl, bk.Description); err != nil {
		return bk, errors.Wrap(err, "inserting book")
	}

	return bk, nil
}

// Delete removes a book by ID from the database.
func (r *RepositoryDb) Delete(bookID string) error {
	if _, err := uuid.Parse(bookID); err != nil {
		return ErrInvalidID
	}

	const q = `
	DELETE FROM
		books
	WHERE
		book_id = $1
	`

	if _, err := r.Db.Exec(q, bookID); err != nil {
		return errors.Wrapf(err, "deleting book with ID %s", bookID)
	}
	return nil
}
