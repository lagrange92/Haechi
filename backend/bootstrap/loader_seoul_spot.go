package bootstrap

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/lagrange92/Haechi/model"
	"github.com/lagrange92/Haechi/utils"
)

func loadSeoulSpots() []model.SeoulSpot {
	csvSpots := loadSpotsCSV()

	seoulSpots := makeSeoulSpots(csvSpots)

	return seoulSpots
}

// loadSpotsCSV reads "./resources/seoul_spot_113.csv" and returns 113 spots data
func loadSpotsCSV() [][]string {
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

func makeSeoulSpots(csvSpots [][]string) []model.SeoulSpot {
	seoulSpots := []model.SeoulSpot{}

	for _, csvSpot := range csvSpots {
		seoulSpots = append(seoulSpots, makeSeoulSpot(csvSpot))
	}

	return seoulSpots
}

func makeSeoulSpot(csvSpot []string) model.SeoulSpot {
	shortCode, err := strconv.Atoi(csvSpot[2])
	if err != nil {
		fmt.Println("Error converting string to int: ", err)
	}

	latFloat, err := strconv.ParseFloat(strings.TrimSpace(csvSpot[4]), 64)
	if err != nil {
		fmt.Println("Error converting string to float: ", err)
	}

	lngFloat, err := strconv.ParseFloat(strings.TrimSpace(csvSpot[5]), 64)
	if err != nil {
		fmt.Println("Error converting string to float: ", err)
	}

	latFloat = utils.TruncateToSixDecimalPlaces(latFloat)
	lngFloat = utils.TruncateToSixDecimalPlaces(lngFloat)

	seoulSpot := model.SeoulSpot{
		Category:  csvSpot[0],
		Code:      csvSpot[1],
		ShortCode: shortCode,
		AreaName:  csvSpot[3],
		Latitude:  latFloat,
		Longitude: lngFloat,
	}

	return seoulSpot
}
