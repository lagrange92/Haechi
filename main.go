package main

import (
	"context"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/tmc/langchaingo/llms/openai"
)

// Handler
func handleHome(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// Echo code snippet
	// e := echo.New()

	// e.GET("/", handleHome)
	// e.Logger.Fatal(e.Start(":1323"))

	// OpenAI code snippet
	llm, err := openai.New()
	if err != nil {
		log.Fatal(err)
	}

	prompt := "What's your model name?"
	completion, err := llm.Call(context.Background(), prompt)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(completion)
}
