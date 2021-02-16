package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pkg/errors"
	"github.com/sophiabrandt/go-rest-api/internal/adapter/database"
	"github.com/sophiabrandt/go-rest-api/internal/env"
	"github.com/sophiabrandt/go-rest-api/internal/server"
	transportHTTP "github.com/sophiabrandt/go-rest-api/internal/transport/http"
)

func main() {
	// make a channel to listen for interrupt or terminal signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())

	log := log.New(os.Stdout, "GO-REST-API: ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	// listen for interrupt signals
	go func() {
		oscall := <-c
		log.Printf("main: system call: %+v", oscall)
		cancel()
	}()

	// run API
	if err := run(ctx, log); err != nil {
		log.Println("main: error:", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, log *log.Logger) error {
	log.Println("main : Started : Application initializing")

	addr := flag.String("addr", "0.0.0.0:4000", "Http network address")
	flag.Parse()

	// databse
	db, err := database.New()
	if err != nil {
		return errors.Wrap(err, "could not start server")
	}
	defer db.Close()

	// initialize global dependencies
	env := env.New(log, db)

	router := transportHTTP.New(env)

	// create server
	srv := server.New(*addr, router)

	go func() {
		log.Printf("main: API listening on %s", *addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("main: %+s\n", err)
		}
	}()

	<-ctx.Done()
	log.Println("main: Start shutdown")

	// shutdown server
	ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err = srv.Shutdown(ctxShutdown); err != nil {
		log.Fatalf("main: Shutdown Failed: %+s", err)
	}

	log.Println("main: API exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}

	return nil
}
