package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/shunsukenagashima/openai-go-sample/pkg/config"
	"github.com/shunsukenagashima/openai-go-sample/pkg/infra/openai"
	"github.com/shunsukenagashima/openai-go-sample/pkg/interface/controller"
	"github.com/shunsukenagashima/openai-go-sample/pkg/usecase"
)

func main() {
	http := &http.Client{}

	cfg, err := config.New()
	if err != nil {
		log.Printf("failed to load config: %v", err)
		os.Exit(1)
	}

	log.Print(cfg)

	oc := openai.NewOpenAIClient(cfg.OpenAIKey, cfg.AIModel, http)
	v := validator.New()
	cu := usecase.NewChatUsecase(oc)
	cc := controller.NewChatController(cu, v)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	r.POST("/chat", cc.SendMessage)

	r.Run()
}
