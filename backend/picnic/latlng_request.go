package picnic

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/lagrange92/Haechi/model"
	"github.com/lagrange92/Haechi/utils"
)

// fillLatLng fills latitude and longitude of picnic spots
func fillLatLng(spot *model.PicnicSpot) {
	lat, lng, _ := getLatLng(spot.Name, os.Getenv("KAKAO_REST_API_KEY"))

	spot.Latitude = lat
	spot.Longitude = lng
}

func getLatLng(keyword, apiKey string) (float64, float64, error) {
	client := &http.Client{}
	url := os.Getenv("KAKAO_KEYWORD_API_URL") + "?query=" + url.QueryEscape(keyword)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, 0, err
	}
	req.Header.Add("Authorization", "KakaoAK "+apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, err
	}

	kwdResp := utils.UnmarshalKaoKwd(body)

	if len(kwdResp.Documents) > 0 {
		kwd := kwdResp.Documents[0]

		lat, _ := strconv.ParseFloat(kwd.Y, 64)
		lng, _ := strconv.ParseFloat(kwd.X, 64)

		return float64(lat), float64(lng), nil
	}

	return 0, 0, fmt.Errorf("No results found")
}
