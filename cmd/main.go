package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "gitlab.com/quickqr/api/docs"
	"gitlab.com/quickqr/api/internal"
)

// @title			Quick QR API
// @version		0.0.0
// @contact.url	https://gitlab.com/quick-qr/api/
// @description.markdown
// @BasePath		/api/
func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	log.Fatal(internal.RunServer(port))

}
