package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/pkg/errors"
	"github.com/sophiabrandt/go-rest-api/internal/server"
)

// App represents the service and contains dependencies,
// e.g., database connection pool, etc.
type App struct{}

func main() {
	log := log.New(os.Stdout, "GO-REST-API: ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	app := App{}
	if err := app.run(log); err != nil {
		log.Println("main: error:", err)
		os.Exit(1)
	}
}

func (app *App) run(log *log.Logger) error {
	log.Println("main : Started : Application initializing")

	addr := flag.String("addr", "0.0.0.0:4000", "Http network address")
	flag.Parse()

	mux := http.NewServeMux()
	srv := server.New(*addr, mux)

	log.Printf("main: API listening on %s", *addr)
	if err := srv.ListenAndServe(); err != nil {
		return errors.Wrap(err, "could not start server")
	}

	return nil
}
