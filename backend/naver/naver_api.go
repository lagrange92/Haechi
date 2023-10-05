package naver

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/lagrange92/Haechi/model"
	"github.com/lagrange92/Haechi/utils"
)

// ReqDir : request direction data to Naver Maps API and return result JSON
func ReqDir(start model.PicnicSpot, goal model.PicnicSpot) (model.NVDirResponse, error) {
	startLng := strconv.FormatFloat(start.Longitude, 'f', 6, 64)
	startLat := strconv.FormatFloat(start.Latitude, 'f', 6, 64)
	goalLng := strconv.FormatFloat(goal.Longitude, 'f', 6, 64)
	goalLat := strconv.FormatFloat(goal.Latitude, 'f', 6, 64)

	startStr := startLng + "," + startLat + ",name=" + start.Name
	goalStr := goalLng + "," + goalLat + ",name=" + goal.Name

	return reqDirInternal(startStr, goalStr)
}

// reqDirInternal : request direction data to Naver Maps API
//
//	curl "https://naveropenapi.apigw.ntruss.com/map-direction/v1/driving?start={출발지}&goal={목적지}&option={탐색옵션}" \
//			-H "X-NCP-APIGW-API-KEY-ID: {애플리케이션 등록 시 발급받은 client id 값}" \
//			-H "X-NCP-APIGW-API-KEY: {애플리케이션 등록 시 발급받은 client secret값}" -v
func reqDirInternal(start string, goal string) (model.NVDirResponse, error) {
	url := "https://naveropenapi.apigw.ntruss.com/map-direction/v1/driving?start=" + start + "&goal=" + goal + "&option=trafast"

	fmt.Println(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("X-NCP-APIGW-API-KEY-ID", "w9sesnkfd4")
	req.Header.Add("X-NCP-APIGW-API-KEY", "44GpgpRwGClMdcAlI2AMsHZq4dSS798f4SaxAYF7")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.NVDirResponse{}, err
	}

	nvDirResponse := utils.UnmarshalNVDir(body)

	// printNVDirResponse(nvDirResponse)

	return nvDirResponse, nil
}

func printNVDirResponse(nvDirResponse model.NVDirResponse) {

	// NVDirResponse 객체 출력
	fmt.Println("Code:", nvDirResponse.Code)
	fmt.Println("Message:", nvDirResponse.Message)
	fmt.Println("CurrentDateTime:", nvDirResponse.CurrentDateTime)

	for _, route := range nvDirResponse.Route.Trafast {
		fmt.Println("Start:", route.Summary.Start.Location)
		fmt.Println("Goal:", route.Summary.Goal.Location)
		fmt.Println("Distance:", route.Summary.Distance)
		fmt.Println("Duration:", route.Summary.Duration)
		fmt.Println("EtaServiceType:", route.Summary.EtaServiceType)
		fmt.Println("Bbox:", route.Summary.Bbox)
		fmt.Println("TollFare:", route.Summary.TollFare)
		fmt.Println("TaxiFare:", route.Summary.TaxiFare)
		fmt.Println("FuelPrice:", route.Summary.FuelPrice)

		for _, path := range route.Path {
			fmt.Print("Path:")
			fmt.Print("Longitude:", path[0])
			fmt.Print("Latitude:", path[1])
			fmt.Println()
		}

		for _, section := range route.Section {
			fmt.Print("Section:")
			fmt.Print("PointIndex:", section.PointIndex)
			fmt.Print("PointCount:", section.PointCount)
			fmt.Print("Distance:", section.Distance)
			fmt.Print("Name:", section.Name)
			fmt.Print("Congestion:", section.Congestion)
			fmt.Print("Speed:", section.Speed)
			fmt.Println()
		}

		for _, guide := range route.Guide {
			fmt.Print("Guide:")
			fmt.Print("PointIndex:", guide.PointIndex)
			fmt.Print("Type:", guide.Type)
			fmt.Print("Instructions:", guide.Instructions)
			fmt.Print("Distance:", guide.Distance)
			fmt.Print("Duration:", guide.Duration)
			fmt.Println()
		}
	}
}
