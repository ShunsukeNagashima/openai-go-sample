package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/shunsukenagashima/openai-go-sample/pkg/domain/usecase"
)

type ChatController struct {
	chatUsecase usecase.ChatUsecase
	validator   *validator.Validate
}

func NewChatController(chatUsecase usecase.ChatUsecase, validator *validator.Validate) *ChatController {
	return &ChatController{
		chatUsecase,
		validator,
	}
}

func (cc *ChatController) SendMessage(ctx *gin.Context) {
	var req struct {
		Message string `json:"message"`
	}

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cc.validator.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := cc.chatUsecase.SendMessage(req.Message)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": resp.Content, "promptTokens": resp.Usage.PromptTokens})
}
