package apiCaller

import (
	"context"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"github.com/yuyuancha/yuyuancha-tool/telegram-bot/config"
	"google.golang.org/api/option"
	"log"
)

const (
	GminiMode = "gemini-1.5-flash"
)

// Gmini Gmini 請求服務
var Gmini = &GminiCaller{}

// GminiCaller Gmini 請求
type GminiCaller struct{}

// RequestTextQuestion 請求文字問題
func (caller *GminiCaller) RequestTextQuestion(question string) (string, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(config.GoogleService.GminiApiKey))
	if err != nil {
		return "", err
	}

	defer func() {
		err = client.Close()
		if err != nil {
			log.Println("關閉 AI client 發生錯誤:", err)
		}
	}()

	model := client.GenerativeModel(GminiMode)
	resp, err := model.GenerateContent(ctx, genai.Text(question))
	if err != nil {
		return "", err
	}

	return caller.parseAIResponse(resp), nil
}

// 解析 AI 回應
func (caller *GminiCaller) parseAIResponse(response *genai.GenerateContentResponse) (results string) {
	for _, candidate := range response.Candidates {
		for _, part := range candidate.Content.Parts {
			results = results + fmt.Sprintf("%v", part)
			fmt.Println(part)
		}
	}
	return
}
