package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/coderflexx/blog-api/internal/database"
	"github.com/coderflexx/blog-api/internal/router"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env")
	}

	// Connect to DB
	database.Connect()

	if os.Getenv("APP_ENV") == "development" {
		database.Seed() // remove this on production.
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on http://localhost:%s \n", port)
	log.Fatal(http.ListenAndServe(":"+port, router.Setup()))
}
