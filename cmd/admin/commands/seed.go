package commands

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/sophiabrandt/go-rest-api/internal/adapter/database"
	"github.com/sophiabrandt/go-rest-api/internal/data/schema"
)

// Seed loads test data into the database.
func Seed() error {
	db, err := database.New()
	if err != nil {
		return errors.Wrap(err, "could not connect to database")
	}
	defer db.Close()

	if err := schema.Seed(db); err != nil {
		return errors.Wrap(err, "seed database")
	}

	fmt.Println("seed data complete")
	return nil
}
