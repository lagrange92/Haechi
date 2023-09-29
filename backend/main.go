package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// SeoulSpot struct from "./resources/seoul_spot_113.csv"
type SeoulSpot struct {
	category  string
	shortCode int
	code      string
	areaName  string
	latitude  float64
	longitude float64
}

type pplChData struct {
	spot SeoulSpot
	data string
}

func truncateToSixDecimalPlaces(f float64) float64 {
	return math.Round(f*1e6) / 1e6
}

func main() {
	var seoulSpots []SeoulSpot
	curPplMap := make(map[SeoulSpot]string)
	var spotPplCh = make(chan pplChData)

	readSpots := readSeoulSpots()

	for _, readSpot := range readSpots {
		shortCode, err := strconv.Atoi(readSpot[2])
		if err != nil {
			fmt.Println("Error converting string to int: ", err)
		}

		latFloat, err := strconv.ParseFloat(strings.TrimSpace(readSpot[4]), 64)
		if err != nil {
			fmt.Println("Error converting string to float: ", err)
		}

		lngFloat, err := strconv.ParseFloat(strings.TrimSpace(readSpot[5]), 64)
		if err != nil {
			fmt.Println("Error converting string to float: ", err)
		}

		latFloat = truncateToSixDecimalPlaces(latFloat)
		lngFloat = truncateToSixDecimalPlaces(lngFloat)

		seoulSpots = append(seoulSpots, SeoulSpot{
			category:  readSpot[0],
			code:      readSpot[1],
			shortCode: shortCode,
			areaName:  readSpot[3],
			latitude:  latFloat,
			longitude: lngFloat,
		})
	}

	for _, seoulSpot := range seoulSpots {
		go getCurPplData(seoulSpot, spotPplCh)
	}

	for i := 0; i < len(seoulSpots); i++ {
		chanData := <-spotPplCh
		curPplMap[chanData.spot] = chanData.data
	}

	// Activate to check if you want to check that data is valid
	savePpl(curPplMap)

	// ====================== Echo code snippet ======================
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
	// ===============================================================
}

func getCurPplData(spot SeoulSpot, ch chan<- pplChData) {
	// Seoul Open API code snippet
	url := "http://openapi.seoul.go.kr:8088/" + os.Getenv("SEOUL_OPEN_API_KEY") + "/json/citydata_ppltn/1/5/" + spot.areaName
	fmt.Println("try ", url)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	ch <- pplChData{spot, string(body)}
}

// readSeoulSpots reads "./resources/seoul_spot_113.csv" and returns 113 spots data
func readSeoulSpots() [][]string {
	f, err := os.Open("./resources/seoul_spot_113.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return records
}

func savePpl(curPplMap map[SeoulSpot]string) {
	// 파일 생성 및 열기
	filename := "./current_people_" + time.Now().Format("060102_1504") + ".txt" // ex) current_people_YYMMDD_HHmm.txt
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Error creating file: ", err)
	}

	defer file.Close()

	num := 1
	for spot, data := range curPplMap {
		_, err := fmt.Fprintf(file, "[%d. %s] %f %f\n%s\n", num, spot.areaName, spot.latitude, spot.longitude, data)
		if err != nil {
			log.Fatal("Error writing to file: ", err)
		}
		num++
	}

	fmt.Println("Data saved to ", filename)
}

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
