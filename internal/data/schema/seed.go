package schema

import (
	"github.com/jmoiron/sqlx"
)

// Seed runs the set of seed-data queries against db. The queries are ran in a
// transaction and rolled back if any fail.
func Seed(db *sqlx.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(seeds); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}

// seeds is a string constant containing all of the queries needed to get the
// db seeded to a useful state for development.
//
// Note that database servers besides PostgreSQL may not support running
// multiple queries as part of the same execution so this single large constant
// may need to be broken up.
const seeds = `
-- Create authors and books
INSERT INTO authors (author_id, name) VALUES
	('bbc79841-7feb-4944-9971-07404558dfdd', 'j.r.r.tolkien'),
	('6ae4a9bf-0bff-40d5-9dbc-ce93819f4208', 'jane austen')
	ON CONFLICT DO NOTHING;

INSERT INTO books (book_id, title, author_id, published_date, image_url, description) VALUES
	('5cf37266-3473-4006-984f-9325122678b7', 'the lord of the rings', 'bbc79841-7feb-4944-9971-07404558dfdd', '1954-07-29', 'https://source.unsplash.com/random/300x300', 'high fantasy novel'),
	('45b5fbd3-755f-4379-8f07-a58d4a30fa2f', 'pride and prejudice', '6ae4a9bf-0bff-40d5-9dbc-ce93819f4208', '1813-01-28', 'https://source.unsplash.com/random/300x300', 'romantic novel of manners')
	ON CONFLICT DO NOTHING;
`

// DeleteAll runs the set of Drop-table queries against db. The queries are ran in a
// transaction and rolled back if any fail.
func DeleteAll(db *sqlx.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(deleteAll); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}

// deleteAll is used to clean the database between tests.
const deleteAll = `
DELETE FROM books;
DELETE FROM authors;
`
