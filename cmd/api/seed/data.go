// we are literally making users out of thin fucking air
package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/twonkista/PikePlace/internal/data"
	"github.com/twonkista/PikePlace/internal/db"
)

func main() {
	_ = godotenv.Load("C:\\Users\\rkoma\\projects\\PikePlace\\.env")

	dbConn, err := db.Open(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	if err := data.CreateUsers(dbConn, "Joji", "kansaidialect6767", 8000000, 2); err != nil {
		log.Fatal(err)
	}

	log.Println("Seed complete")
}
