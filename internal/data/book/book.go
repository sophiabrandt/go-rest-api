package book

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// BookRepositoryDb defines the repository for the book service.
type RepositoryDb struct {
	DB *sqlx.DB
}

// Info is the book model.
type Info struct {
	ID            string `db:"book_id" json:"book_id"`
	AuthorID      string `db:"author_id" json:"author_id,omitempty"`
	AuthorName    string `db:"author_name" json:"author_name"`
	Title         string `db:"title" json:"title"`
	PublishedDate string `db:"published_date" json:"published_date"`
	ImageUrl      string `db:"image_url" json:"image_url"`
	Description   string `db:"description" json:"description"`
}

// Repo is the interface for the book repository.
type Repo interface {
	Query() ([]Info, error)
}

// New returns a pointer to a book repo.
func New(db *sqlx.DB) *RepositoryDb {
	return &RepositoryDb{DB: db}
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
	if err := r.DB.Select(&books, q); err != nil {
		return nil, errors.Wrap(err, "selecting books")
	}
	return books, nil
}
