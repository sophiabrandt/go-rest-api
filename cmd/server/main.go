package main

import (
	"flag"
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/sophiabrandt/go-rest-api/internal/env"
	"github.com/sophiabrandt/go-rest-api/internal/server"
	transportHTTP "github.com/sophiabrandt/go-rest-api/internal/transport/http"
)

func main() {
	log := log.New(os.Stdout, "GO-REST-API: ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	if err := run(log); err != nil {
		log.Println("main: error:", err)
		os.Exit(1)
	}
}

func run(log *log.Logger) error {
	log.Println("main : Started : Application initializing")

	addr := flag.String("addr", "0.0.0.0:4000", "Http network address")
	flag.Parse()

	// initialize gloabl dependencies
	env := env.New(log)

	router := transportHTTP.New(env)

	// create server
	srv := server.New(*addr, router)

	log.Printf("main: API listening on %s", *addr)
	if err := srv.ListenAndServe(); err != nil {
		return errors.Wrap(err, "could not start server")
	}

	return nil
}
