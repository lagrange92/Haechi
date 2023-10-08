package chat

import (
	"context"
	"log"
	"strconv"
	"strings"

	"github.com/lagrange92/Haechi/model"
	"github.com/lagrange92/Haechi/picnic"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

var gptLLM *openai.LLM
var gptCallOptions []llms.CallOption

// Chat : chat with OpenAI ChatGPT
func Chat(prompt string) (model.ChatResponseData, error) {
	if gptLLM == nil {
		llm, err := openai.New()
		if err != nil {
			log.Fatal(err)
		}

		options := []llms.CallOption{
			llms.WithMaxTokens(1024),
		}

		gptLLM = llm
		gptCallOptions = options
	}

	reqPrompt := genReqPrompt(prompt)

	llmResponse, err := gptLLM.Call(context.Background(), reqPrompt, gptCallOptions...)
	if err != nil {
		return model.ChatResponseData{}, err
	}

	// llmResponse := "start->어린이대공원역\nlayovers->신사역, 신논현역\ngoal->신림역\nppl->0.5\n"

	picnicReq := picnic.ConvertLLMResp(llmResponse)

	picnicPlan, _ := picnic.MakePlan(picnicReq)

	picnicPlanResp := genPicnicPlanResp(picnicPlan)

	return picnicPlanResp, nil
}

func genPicnicPlanResp(plan model.PicnicResponse) model.ChatResponseData {
	chatData := strings.Builder{}
	chatData.WriteString("출발지는 ")
	chatData.WriteString(plan.Start.Name)
	// chatData.WriteString("(")
	// chatData.WriteString(strconv.FormatFloat(plan.Start.Latitude, 'f', 6, 64))
	// chatData.WriteString(", ")
	// chatData.WriteString(strconv.FormatFloat(plan.Start.Longitude, 'f', 6, 64))
	// chatData.WriteString(")")
	chatData.WriteString(", ")

	if len(plan.LayOvers) > 0 {
		chatData.WriteString("경유지는 ")
		first := true
		for _, layOver := range plan.LayOvers {
			if first == false {
				chatData.WriteString(", ")
			}
			chatData.WriteString(layOver.Name)
			// chatData.WriteString("(")
			// chatData.WriteString(strconv.FormatFloat(layOver.Latitude, 'f', 6, 64))
			// chatData.WriteString(", ")
			// chatData.WriteString(strconv.FormatFloat(layOver.Longitude, 'f', 6, 64))
			// chatData.WriteString(")")

			first = false
		}
		chatData.WriteString("이고, ")
	}
	chatData.WriteString("목적지는 ")
	chatData.WriteString(plan.Goal.Name)
	// chatData.WriteString("(")
	// chatData.WriteString(strconv.FormatFloat(plan.Goal.Latitude, 'f', 6, 64))
	// chatData.WriteString(", ")
	// chatData.WriteString(strconv.FormatFloat(plan.Goal.Longitude, 'f', 6, 64))
	// chatData.WriteString(")")
	chatData.WriteString("입니다.")

	suggestData := strings.Builder{}
	suggestData.WriteString("[{\"name\": \"" + plan.Start.Name + "\", \"lat\": " + strconv.FormatFloat(plan.Start.Latitude, 'f', 6, 64) + ", \"lng\": " + strconv.FormatFloat(plan.Start.Longitude, 'f', 6, 64) + "}, ")
	for _, sugSpot := range plan.LayOvers {
		suggestData.WriteString("{\"name\": \"" + sugSpot.Name + "\", \"lat\": " + strconv.FormatFloat(sugSpot.Latitude, 'f', 6, 64) + ", \"lng\": " + strconv.FormatFloat(sugSpot.Longitude, 'f', 6, 64) + "}, ")
	}
	suggestData.WriteString("{\"name\": \"" + plan.Goal.Name + "\", \"lat\": " + strconv.FormatFloat(plan.Goal.Latitude, 'f', 6, 64) + ", \"lng\": " + strconv.FormatFloat(plan.Goal.Longitude, 'f', 6, 64) + "}]")

	var resp model.ChatResponseData

	resp.Chat = chatData.String()
	resp.Suggest = suggestData.String()

	return resp
}
