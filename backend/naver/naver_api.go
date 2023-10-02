package naver

import (
	"fmt"
	"net/http"
)

// RequestDirection : request direction data to Naver Maps API
//
//	curl "https://naveropenapi.apigw.ntruss.com/map-direction/v1/driving?start={출발지}&goal={목적지}&option={탐색옵션}" \
//			-H "X-NCP-APIGW-API-KEY-ID: {애플리케이션 등록 시 발급받은 client id 값}" \
//			-H "X-NCP-APIGW-API-KEY: {애플리케이션 등록 시 발급받은 client secret값}" -v
func RequestDirection(start string, goal string) []string {
	url := "https://naveropenapi.apigw.ntruss.com/map-direction/v1/driving?start=" + start + "&goal=" + goal + "&option=trafast"

	// fmt.Println(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("X-NCP-APIGW-API-KEY-ID", "w9sesnkfd4")
	req.Header.Add("X-NCP-APIGW-KEY", "44GpgpRwGClMdcAlI2AMsHZq4dSS798f4SaxAYF7")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)

	return []string{}
}
