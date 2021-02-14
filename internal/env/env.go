package env

import "github.com/dimfeld/httptreemux/v5"

// Env defines the local app context and holds global
// dependencies
type Env struct {
	Router *httptreemux.ContextMux
}

// New creates a new pointer to an Env struct.
func New() *Env {
	router := httptreemux.NewContextMux()
	return &Env{
		Router: router,
	}
}
