package main

import (
	"ethFees/internal/api"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server := api.NewServer()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(server.Start(":" + port))
}
