package env

import (
	"log"

	"github.com/dimfeld/httptreemux/v5"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
)

// Env defines the local app context and holds global
// dependencies.
type Env struct {
	Log       *log.Logger
	Db        *sqlx.DB
	Router    *httptreemux.ContextMux
	Validator *validator.Validate
}

// New creates a new pointer to an Env struct.
func New(log *log.Logger, db *sqlx.DB, validator *validator.Validate) *Env {
	router := httptreemux.NewContextMux()
	return &Env{
		Log:       log,
		Db:        db,
		Router:    router,
		Validator: validator,
	}
}
