package book

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// BookRepositoryDb defines the repository for the book service.
type RepositoryDb struct {
	Db *sqlx.DB
}

// Repo is the interface for the book repository.
type Repo interface {
	Query() ([]Info, error)
	Create(book NewBook) (Info, error)
}

// New returns a pointer to a book repo.
func New(db *sqlx.DB) *RepositoryDb {
	return &RepositoryDb{Db: db}
}

// Query retrieves all books from the database.
func (r *RepositoryDb) Query() ([]Info, error) {
	const q = `
	SELECT
		b.book_id, b.title, b.published_date, b.image_url, b.description,
		a.name AS author_name
	FROM books as b
	LEFT JOIN
		authors AS a ON b.author_id = a.author_id
	ORDER BY b.book_id
	`
	books := []Info{}
	if err := r.Db.Select(&books, q); err != nil {
		return nil, errors.Wrap(err, "selecting books")
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

	// create new book model for the database
	bk := Info{
		ID:            uuid.New().String(),
		Title:         book.Title,
		AuthorID:      author_id,
		AuthorName:    book.AuthorName,
		PublishedDate: book.PublishedDate,
		ImageUrl:      book.ImageUrl,
		Description:   book.Description,
	}

	const q = `
	INSERT INTO books
		(book_id, title, author_id, published_date, image_url, description)
	VALUES
		($1, $2, $3, $4, $5, $6)`

	if _, err := r.Db.Exec(q, bk.ID, bk.Title, bk.AuthorID, bk.PublishedDate, bk.ImageUrl, bk.Description); err != nil {
		return Info{}, errors.Wrap(err, "inserting book")
	}

	return bk, nil
}
