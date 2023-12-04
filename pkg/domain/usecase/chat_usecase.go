package usecase

import "github.com/shunsukenagashima/openai-go-sample/pkg/infra/openai"

type ChatUsecase interface {
	SendMessage(message string) (*openai.ChatResponse, error)
}
