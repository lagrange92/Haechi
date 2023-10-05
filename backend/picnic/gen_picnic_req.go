package picnic

import (
	"strings"

	"github.com/lagrange92/Haechi/model"
)

// ConvertLLMResp : generate picnic request from llm response
func ConvertLLMResp(llmResponse string) model.PicnicRequest {
	picnicReq := model.PicnicRequest{
		Start:    model.PicnicSpot{},
		LayOvers: []model.PicnicSpot{},
		Goal:     model.PicnicSpot{},
	}

	lines := strings.Split(llmResponse, "\n")

	for _, line := range lines {
		tokens := strings.Split(line, "->")
		if tokens[0] == "start" {
			picnicReq.Start.Name = tokens[1]
			fillLatLng(&picnicReq.Start)

		} else if tokens[0] == "layovers" {
			layOvers := strings.Split(tokens[1], ",")
			for _, layOver := range layOvers {
				layOver := model.PicnicSpot{
					Name: layOver,
				}
				fillLatLng(&layOver)
				picnicReq.LayOvers = append(picnicReq.LayOvers, layOver)
			}
		} else if tokens[0] == "goal" {
			picnicReq.Goal.Name = tokens[1]
			fillLatLng(&picnicReq.Goal)
		} else if tokens[0] == "ppl" {
			// continue
		}
	}

	return picnicReq
}
