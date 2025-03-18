package apiCaller

import (
	"context"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"github.com/yuyuancha-tool/pokemon-card-tool/config"
	"google.golang.org/api/option"
	"log"
	"os"
)

const (
	GminiMode = "gemini-1.5-flash"
)

// GminiCaller Gmini 請求
type GminiCaller struct{}

// NewGminiCaller new Gmini caller
func NewGminiCaller() *GminiCaller {
	return &GminiCaller{}
}

// RequestTextQuestionByImage 透過圖片請求文字問題
func (caller *GminiCaller) RequestTextQuestionByImage(question string, pathList ...string) (string, error) {
	images := make([][]byte, 0, len(pathList))
	for _, path := range pathList {
		imageData, err := os.ReadFile(path)
		if err != nil {
			return "", err
		}
		images = append(images, imageData)
	}

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

	var prompt []genai.Part
	for _, image := range images {
		prompt = append(prompt, genai.ImageData("png", image))
	}
	prompt = append(prompt, genai.Text(question))

	resp, err := model.GenerateContent(ctx, prompt...)
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
			if config.GoogleService.HasGminiApiParseLog {
				fmt.Println(part)
			}
		}
	}
	return
}
