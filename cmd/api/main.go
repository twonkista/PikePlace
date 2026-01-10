package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/twonkista/PikePlace/internal/db"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found; relying on environment variables")
	}
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is empty (not loaded from .env and not set in environment)")
	}

	dbConn, err := db.Open(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("db open failed: %v", err)
	}
	log.Println("Connected to database.")

	cfg := config{
		addr: ":8080",
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
