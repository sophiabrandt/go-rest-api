// Package schema contains the database schema, migrations and seeding data.
package schema

import (
	"github.com/dimiro1/darwin"
	"github.com/jmoiron/sqlx"
)

// Migrate attempts to bring the schema for db up to date with the migrations
// defined in this package.
func Migrate(db *sqlx.DB) error {
	driver := darwin.NewGenericDriver(db.DB, darwin.SqliteDialect{})
	d := darwin.New(driver, migrations, nil)
	return d.Migrate()
}

// migrations contains the queries needed to construct the database schema.
// Entries should never be removed once they have been run in production.
//
// Using constants in a .go file is an easy way to ensure the schema is part
// of the compiled executable and avoids pathing issues with the working
// directory. It has the downside that it lacks syntax highlighting and may be
// harder to read for some cases compared to using .sql files. You may also
// consider a combined approach using a tool like packr or go-bindata.
var migrations = []darwin.Migration{
	{
		Version:     1.1,
		Description: "Create table authors",
		Script: `
CREATE TABLE authors (
	author_id       UUID,
	name            TEXT NOT NULL,
	PRIMARY KEY (author_id)
);`,
	},
	{
		Version:     1.2,
		Description: "Create table books",
		Script: `
CREATE TABLE books (
	book_id         UUID,
	author_id       UUID,
	title           TEXT NOT NULL,
	published_date  TIMESTAMP,
	image_url       TEXT,
	description     TEXT,
	PRIMARY KEY (book_id),
	FOREIGN KEY (author_id) REFERENCES authors(author_id) ON DELETE CASCADE
);`,
	},
}
