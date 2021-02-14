package http

import (
	"net/http"

	"github.com/sophiabrandt/go-rest-api/internal/env"
)

// Handler takes a configured Env.
type Handler struct {
	E *env.Env
	H func(E *env.Env, w http.ResponseWriter, r *http.Request) error
}

// ServeHTTP allows the Handler to satisy the http.Handler interface.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.H(h.E, w, r)
	if err != nil {
		h.E.Router.ServeHTTP(w, r)
	}
}
