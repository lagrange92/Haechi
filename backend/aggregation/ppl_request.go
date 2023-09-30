package aggreation

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/lagrange92/Haechi/model"
	"github.com/lagrange92/Haechi/utils"
)

// GetCurPplData gets current population data from Seoul Open API
func requestPpl(spot model.SeoulSpot, ch chan<- model.PpltnData) {
	url := model.SeoulBaseURL + spot.AreaName
	body := requestPplOnURL(spot, url)

	ppl := makePpl(body, spot)

	ch <- ppl
}

func requestPplOnURL(spot model.SeoulSpot, url string) []byte {
	fmt.Println("try aggregation to ", url, "...")

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body
}

func makePpl(body []byte, spot model.SeoulSpot) model.PpltnData {
	pplJSON := utils.Unmarshal(body)

	if len(pplJSON.Ppltn) == 0 {
		fmt.Println("No pplJSON data for ", spot.AreaName)
		return model.PpltnData{}
	}

	ppl := utils.Convert(pplJSON)

	ppl.AreaLatitude = fmt.Sprintf("%f", spot.Latitude)
	ppl.AreaLongitude = fmt.Sprintf("%f", spot.Longitude)

	return ppl
}
