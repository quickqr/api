package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gitlab.com/quick-qr/server/internal"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	log.Fatal(internal.RunServer(port))

}
