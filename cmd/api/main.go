package main

import (
	"go-api/internal/env"
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

	cfg := config{
		addr: env.Getstring("ADDR", addr),
	}
	app := &application{
		config: cfg,
	}

	app.run(app.mount())
}
