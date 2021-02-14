package main

import (
	"log"
	"os"
)

// App represents the service and contains dependencies,
// e.g., database connection pool, etc.
type App struct {}

func main() {
	log := log.New(os.Stdout, "GO-REST-API: ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	if err := run(log); err != nil {
		log.Println("main: error:", err)
		os.Exit(1)
	}
}

func run(log *log.Logger) error {
	log.Println("main : Started : Application initializing")
	
	return nil
}
