package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lagrange92/Haechi/bootstrap"
	"github.com/lagrange92/Haechi/store"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	store.SeoulBaseURL = os.Getenv("SEOUL_OPEN_API_BASE_URL") +
		os.Getenv("SEOUL_OPEN_API_KEY") +
		os.Getenv("SEOUL_OPEN_API_URL_SUFFIX")
}

func main() {
	bootstrap.Bootstrap()

	go bootstrap.ActivateWorker()

	server := CreateServer()

	err := StartServer(server)
	if err != nil {
		log.Fatal(err)
	}
}
