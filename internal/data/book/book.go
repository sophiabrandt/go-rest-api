package book

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// BookRepositoryDb defines the repository for the book service.
type RepositoryDb struct {
	DB *sqlx.DB
}

// Info is the book model.
type Info struct {
	ID            string    `db:"book_id" json:"book_id"`
	Title         string    `db:"title" json:"title"`
	Author        string    `db:"author" json:"book"`
	PublishedDate time.Time `db:published_date" json:"published_date"`
	ImageUrl      string    `db:"image_url" json:"image_url"`
	Description   string    `db:"description" json:"description"`
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
	SELECT b.*
	FROM books as b
	ORDER BY book_id
	`
	books := []Info{}
	if err := r.DB.Select(&books, q); err != nil {
		return nil, errors.Wrap(err, "selecting books")
	}
	return books, nil
}
