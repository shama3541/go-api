package main

import (
	"go-api/internal/env"
	"go-api/internal/store"
	"log"
	"os"

	"github.com/joho/godotenv"
	// Import the env package
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Access env variable
	addr := os.Getenv("ADDR")

	store := store.NewPostgresDb(nil)
	cfg := config{
		addr: env.Getstring("ADDR", addr),
	}
	app := &application{
		config: cfg,
		store:  store,
	}

	app.run(app.mount())
}
