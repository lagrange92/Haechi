package chat

import (
	"context"
	"log"

	"github.com/tmc/langchaingo/llms/openai"
)

// Chat : chat with OpenAI ChatGPT
func Chat(prompt string) (string, error) {
	llm, err := openai.New()
	if err != nil {
		log.Fatal(err)
	}

	completion, err := llm.Call(context.Background(), prompt)
	if err != nil {
		return "", err
	}

	return completion, nil
}
