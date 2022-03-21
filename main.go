package main

import (
	"log"

	"github.com/almcr/crud-go/app"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	server := app.NewServer()
	server.Run()
}
