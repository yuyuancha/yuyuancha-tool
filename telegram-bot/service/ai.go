package service

import (
	"context"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"github.com/yuyuancha/yuyuancha-tool/telegram-bot/config"
	"google.golang.org/api/option"
	"log"
)

const (
	AIModel = "gemini-1.5-flash"
)

// 請求 AI 服務
func requestAI(question string) string {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(config.GoogleService.GminiApiKey))
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = client.Close()
		if err != nil {
			log.Println("關閉 AI client 發生錯誤:", err)
		}
	}()

	model := client.GenerativeModel(AIModel)
	resp, err := model.GenerateContent(ctx, genai.Text(question))
	if err != nil {
		log.Fatal(err)
	}

	return parseAIResponse(resp)
}

// 解析 AI 回應
func parseAIResponse(response *genai.GenerateContentResponse) (results string) {
	for _, candidate := range response.Candidates {
		for _, part := range candidate.Content.Parts {
			results = results + fmt.Sprintf("%v", part)
			fmt.Println(part)
		}
	}
	return
}
