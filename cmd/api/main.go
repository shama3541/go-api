package main

import (
	"go-api/internal/db"
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

	cfg := config{
		addr: env.Getstring("ADDR", addr),
		DB: Dbconfig{
			addr:         env.Getstring("DB_ADDR", "postgres://user:adminpassword@localhost:5432/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.Getstring("DB_MAX_IDLE_TIME", "15min"),
		},
	}
	db, err := db.New(cfg.DB.addr, cfg.DB.maxOpenConns, cfg.DB.maxIdleConns, cfg.DB.maxIdleTime)
	if err != nil {
		log.Printf("Error while connecting to db:%v", err)
	}
	store := store.NewPostgresDb(db)
	app := &application{
		config: cfg,
		store:  store,
	}

	app.run(app.mount())
}
