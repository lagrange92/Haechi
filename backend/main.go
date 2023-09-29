package main

import (
	"encoding/csv"
	"encoding/json"
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
	"github.com/rs/cors"
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

var seoulSpots []SeoulSpot
var curPplMap = make(map[SeoulSpot]string)
var spotPplCh = make(chan pplChData)

func truncateToSixDecimalPlaces(f float64) float64 {
	return math.Round(f*1e6) / 1e6
}

type FcstPpltnJSON struct {
	FcstTime       string `json:"FCST_TIME"`
	FcstCongestLvl string `json:"FCST_CONGEST_LVL"`
	FcstPpltnMin   string `json:"FCST_PPLTN_MIN"`
	FcstPpltnMax   string `json:"FCST_PPLTN_MAX"`
}

type PpltnJSON struct {
	AreaName          string          `json:"AREA_NM"`
	AreaCode          string          `json:"AREA_CD"`
	AreaCongestLvl    string          `json:"AREA_CONGEST_LVL"`
	AreaCongestMsg    string          `json:"AREA_CONGEST_MSG"`
	AreaPpltnMin      string          `json:"AREA_PPLTN_MIN"`
	AreaPpltnMax      string          `json:"AREA_PPLTN_MAX"`
	MalePpltnRate     string          `json:"MALE_PPLTN_RATE"`
	FemalePpltnRate   string          `json:"FEMALE_PPLTN_RATE"`
	PpltnRate0        string          `json:"PPLTN_RATE_0"`
	PpltnRate10       string          `json:"PPLTN_RATE_10"`
	PpltnRate20       string          `json:"PPLTN_RATE_20"`
	PpltnRate30       string          `json:"PPLTN_RATE_30"`
	PpltnRate40       string          `json:"PPLTN_RATE_40"`
	PpltnRate50       string          `json:"PPLTN_RATE_50"`
	PpltnRate60       string          `json:"PPLTN_RATE_60"`
	PpltnRate70       string          `json:"PPLTN_RATE_70"`
	ResntPpltnRate    string          `json:"RESNT_PPLTN_RATE"`
	NonResntPpltnRate string          `json:"NON_RESNT_PPLTN_RATE"`
	ReplaceYn         string          `json:"REPLACE_YN"`
	PpltnTime         string          `json:"PPLTN_TIME"`
	FcstYn            string          `json:"FCST_YN"`
	FcstPpltn         []FcstPpltnJSON `json:"FCST_PPLTN"`
}

type SeoulCityJSON struct {
	Ppltn []PpltnJSON `json:"SeoulRtd.citydata_ppltn"`
}

type PpltnData struct {
	AreaName      string `json:"areaName"`
	AreaCode      string `json:"areaCode"`
	AreaLongitude string `json:"areaLongitude"`
	AreaLatitude  string `json:"areaLatitude"`
	AreaAvgPpltn  string `json:"areaAvgPpltn"`
}

// Handler
func handlePpltn(c echo.Context) error {

	var recvSeoulCity SeoulCityJSON

	var sendPpltns []PpltnData

	for spot, data := range curPplMap {
		err := json.Unmarshal([]byte(data), &recvSeoulCity)
		if err != nil {
			fmt.Println(err)
			continue
		}

		var sendPpltn PpltnData
		sendPpltn.AreaName = recvSeoulCity.Ppltn[0].AreaName
		sendPpltn.AreaCode = recvSeoulCity.Ppltn[0].AreaCode
		sendPpltn.AreaLatitude = fmt.Sprintf("%f", spot.latitude)
		sendPpltn.AreaLongitude = fmt.Sprintf("%f", spot.longitude)

		ppltnMin, _ := strconv.Atoi(recvSeoulCity.Ppltn[0].AreaPpltnMin)
		ppltnMax, _ := strconv.Atoi(recvSeoulCity.Ppltn[0].AreaPpltnMax)
		sendPpltn.AreaAvgPpltn = fmt.Sprintf("%d", (ppltnMin+ppltnMax)/2)

		sendPpltns = append(sendPpltns, sendPpltn)
	}

	// normalize average population to 0~1
	normalMin := 987654321
	normalMax := 0

	for _, sendPpltn := range sendPpltns {
		avgPpltn, _ := strconv.Atoi(sendPpltn.AreaAvgPpltn)
		if avgPpltn < normalMin {
			normalMin = avgPpltn
		}
		if avgPpltn > normalMax {
			normalMax = avgPpltn
		}
	}

	for i, sendPpltn := range sendPpltns {
		avgPpltn, _ := strconv.Atoi(sendPpltn.AreaAvgPpltn)
		normalPpltn := float64(avgPpltn-normalMin) / float64(normalMax-normalMin)
		sendPpltns[i].AreaAvgPpltn = fmt.Sprintf("%f", normalPpltn)
	}

	return c.JSON(http.StatusOK, sendPpltns)
}

func handleMapData(c echo.Context) error {
	// return seoulSpots data as json
	fmt.Println(seoulSpots)
	return c.JSON(http.StatusOK, seoulSpots)
}

func main() {

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
	// savePpl(curPplMap)

	// ====================== Echo code snippet ======================
	e := echo.New()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
	})

	e.Use(echo.WrapMiddleware(c.Handler))

	e.GET("/", handleHome)
	e.GET("/ppltn", handlePpltn)
	e.Logger.Fatal(e.Start(":1323"))
	// ===============================================================

	// =================== OpenAI code snippet =======================
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
