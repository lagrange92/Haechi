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
func Chat(prompt string) (string, error) {
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
		return "", err
	}

	picnicReq := picnic.ConvertLLMResp(llmResponse)

	picnicPlan, _ := picnic.MakePlan(picnicReq)

	picnicPlanStr := genPicnicPlanStr(picnicPlan)

	return picnicPlanStr, nil
}

func genPicnicPlanStr(plan model.PicnicResponse) string {
	picnicPlanStr := strings.Builder{}
	picnicPlanStr.WriteString("출발지는 ")
	picnicPlanStr.WriteString(plan.Start.Name)
	picnicPlanStr.WriteString("(")
	picnicPlanStr.WriteString(strconv.FormatFloat(plan.Start.Latitude, 'f', 6, 64))
	picnicPlanStr.WriteString(", ")
	picnicPlanStr.WriteString(strconv.FormatFloat(plan.Start.Longitude, 'f', 6, 64))
	picnicPlanStr.WriteString(")")
	picnicPlanStr.WriteString(", ")

	if len(plan.LayOvers) > 0 {
		picnicPlanStr.WriteString("경유지는 ")
		first := true
		for _, layOver := range plan.LayOvers {
			if first == false {
				picnicPlanStr.WriteString(", ")
			}
			picnicPlanStr.WriteString(layOver.Name)
			picnicPlanStr.WriteString("(")
			picnicPlanStr.WriteString(strconv.FormatFloat(layOver.Latitude, 'f', 6, 64))
			picnicPlanStr.WriteString(", ")
			picnicPlanStr.WriteString(strconv.FormatFloat(layOver.Longitude, 'f', 6, 64))
			picnicPlanStr.WriteString(")")

			first = false
		}
		picnicPlanStr.WriteString("이고, ")
	}
	picnicPlanStr.WriteString("목적지는 ")
	picnicPlanStr.WriteString(plan.Goal.Name)
	picnicPlanStr.WriteString("(")
	picnicPlanStr.WriteString(strconv.FormatFloat(plan.Goal.Latitude, 'f', 6, 64))
	picnicPlanStr.WriteString(", ")
	picnicPlanStr.WriteString(strconv.FormatFloat(plan.Goal.Longitude, 'f', 6, 64))
	picnicPlanStr.WriteString(")")
	picnicPlanStr.WriteString("입니다.")

	return picnicPlanStr.String()
}
