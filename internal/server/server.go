package server

import (
	"net/http"
	"time"

	"github.com/zenazn/goji/graceful"
)

// New creates a new graceful server.
func New(address string, handler http.Handler) *graceful.Server {
	srv := graceful.Server{
		Addr:         address,
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Minute,
	}
	return &srv
}
