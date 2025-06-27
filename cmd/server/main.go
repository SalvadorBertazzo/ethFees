package main

import (
	"ethFees/internal/api"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server := api.NewServer()

	log.Fatal(server.Start(":8080"))
}
