package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
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
	// llm, err := openai.New()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// prompt := "What's your model name?"
	// completion, err := llm.Call(context.Background(), prompt)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(completion)

	// Seoul Open API code snippet
	// url := "http://openapi.seoul.go.kr:8088/" + os.Getenv("SEOUL_OPEN_API_KEY") + "/json/citydata_ppltn/1/5/광화문·덕수궁"

	// resp, err := http.Get(url)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer resp.Body.Close()

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(body))
}
