package env

import (
	"log"

	"github.com/dimfeld/httptreemux/v5"
	"github.com/jmoiron/sqlx"
)

// Env defines the local app context and holds global
// dependencies.
type Env struct {
	Log    *log.Logger
	DB     *sqlx.DB
	Router *httptreemux.ContextMux
}

// New creates a new pointer to an Env struct.
func New(log *log.Logger, db *sqlx.DB) *Env {
	router := httptreemux.NewContextMux()
	return &Env{
		Log:    log,
		DB:     db,
		Router: router,
	}
}
