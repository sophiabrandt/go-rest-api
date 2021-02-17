// https://blog.questionable.services/article/http-handler-error-handling-revisited/
package web

import (
	"encoding/json"
	"net/http"

	"github.com/sophiabrandt/go-rest-api/internal/env"
)

// Error represents a handler error. It provides methods for a HTTP status
// code and embeds the built-in error interface.
type Error interface {
	error
	Status() int
}

// ErrorResponse is the client response struct for errors.
type ErrorResponse struct {
	Error  string   `json:"error"`
	Fields []string `json:"fields,omitempty"`
}

// StatusError represents an error with an associated HTTP status code.
type StatusError struct {
	Err  error
	Code int
}

// Allows StatusError to satisfy the error interface.
func (se StatusError) Error() string {
	return se.Err.Error()
}

// Returns our HTTP status code.
func (se StatusError) Status() int {
	return se.Code
}

// handler takes a configured Env.
type handler struct {
	E *env.Env
	H func(E *env.Env, w http.ResponseWriter, r *http.Request) error
}

// ServeHTTP allows the Handler to satisy the http.Handler interface.
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.H(h.E, w, r)
	if err != nil {
		switch e := err.(type) {
		case Error:
			// We can retrieve the status here and write out a specific
			// HTTP status code.
			h.E.Log.Printf("HTTP %d - %s", e.Status(), e)
			response := ErrorResponse{e.Error(), nil}
			respond(h.E, w, response, e.Status())
		default:
			// Any error types we don't specifically look out for default
			// to serving a HTTP 500
			response := ErrorResponse{http.StatusText(http.StatusInternalServerError), nil}
			respond(h.E, w, response, http.StatusInternalServerError)
		}
	}
}

// use wraps middleware around handlers.
func use(h handler, middleware ...func(http.Handler) http.Handler) http.Handler {
	var res http.Handler = h
	for _, m := range middleware {
		res = m(res)
	}

	return res
}

// respond answers the client with JSON.
func respond(e *env.Env, w http.ResponseWriter, data interface{}, statusCode int) error {
	// If there is nothing to marshal then set status code and return.
	if statusCode == http.StatusNoContent {
		w.WriteHeader(statusCode)
		return nil
	}

	// Convert the response value to JSON.
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Set the content type and headers once we know marshaling has succeeded.
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	w.Header().Set("X-Frame-Options", "deny")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	// Write the status code to the response.
	w.WriteHeader(statusCode)

	// Send the result back to the client.
	if _, err := w.Write(jsonData); err != nil {
		return err
	}

	return nil
}
