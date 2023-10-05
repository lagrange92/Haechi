package chat

import "strings"

func genReqPrompt(prompt string) string {
	reqPromptBuilder := strings.Builder{}
	reqPromptBuilder.WriteString("\"" + prompt + "\"")
	reqPromptBuilder.WriteString(
		`위 문장에서 출발지, 경유지, 도착지, 인구밀도를 아래 형식으로 대답해줘.\
		다른 말은 하나도 하지 말고 Note 등의 코멘트도 하나도 하지 마.\

		start->출발지 이름\
		layovers->경유지: 경유지 이름 목록(csv형식)\
		goal->도착지: 도착지 이름\
		ppl->인구밀도: 숫자 (0~1 사이)`)

	return reqPromptBuilder.String()
}
