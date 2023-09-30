package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lagrange92/Haechi/bootstrap"
	"github.com/lagrange92/Haechi/model"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	model.SeoulBaseURL = "http://openapi.seoul.go.kr:8088/" + os.Getenv("SEOUL_OPEN_API_KEY") + "/json/citydata_ppltn/1/5/"
}

func main() {
	bootstrap.Load()

	CreateServer()

	err := StartServer()
	if err != nil {
		log.Fatal(err)
	}
}
