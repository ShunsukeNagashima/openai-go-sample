package usecase

import "github.com/shunsukenagashima/openai-go-sample/pkg/infra/openai"

type ChatUsecase struct {
	openaiClient openai.OpenAICommunicator
}

func NewChatUsecase(openaiClient openai.OpenAICommunicator) *ChatUsecase {
	return &ChatUsecase{
		openaiClient: openaiClient,
	}
}

func (u *ChatUsecase) SendMessage(message string) (*openai.ChatResponse, error) {
	userPrompt := &openai.Prompt{
		Role:    openai.User,
		Content: message,
	}

	systemPrompt := &openai.Prompt{
		Role:    openai.System,
		Content: "あなたは関西人です。コテコテの関西弁で話してください。",
	}

	var prompts []*openai.Prompt
	prompts = append(prompts, userPrompt)
	prompts = append(prompts, systemPrompt)

	resp, err := u.openaiClient.SendMessage(prompts)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
