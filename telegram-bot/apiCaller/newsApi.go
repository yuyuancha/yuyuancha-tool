package apiCaller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/yuyuancha/yuyuancha-tool/telegram-bot/config"
	"github.com/yuyuancha/yuyuancha-tool/telegram-bot/model"
	"io"
	"log"
	"net/http"
)

const (
	NewsApiDailyTopHeadlinesURL = "https://newsapi.org/v2/top-headlines"
)

var News = NewsApiCaller{}

// NewsApiCaller 請求新聞結構
type NewsApiCaller struct{}

// GetDailyTopHeadlines 取得每日頭條
func (caller *NewsApiCaller) GetDailyTopHeadlines() ([]*model.NewsArticles, error) {
	url := fmt.Sprintf("%s?sources=%s&sortBy=popularity&pageSize=5&apiKey=%s", NewsApiDailyTopHeadlinesURL, "bbc-news", config.NewsApi.ApiKey)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = response.Body.Close()
		if err != nil {
			log.Println("關閉 News Api response 發生錯誤:", err)
		}
	}()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("請求 News Api status code error: " + response.Status)
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var results NewsApiDailyTopHeadlinesResponse

	err = json.Unmarshal(bodyBytes, &results)
	if err != nil {
		return nil, err
	}

	return results.ToModel(), nil
}
