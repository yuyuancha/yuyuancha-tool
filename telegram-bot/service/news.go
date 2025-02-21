package service

import (
	"fmt"
	"github.com/yuyuancha/yuyuancha-tool/telegram-bot/webCrawler"
)

// 取得新聞頭條
func getDailyTopHeadlines() string {
	//results, err := apiCaller.News.GetDailyTopHeadlines()
	//if err != nil {
	//	panic(err)
	//}
	results := webCrawler.GetDailyTopHeadlines()

	var str string
	for _, result := range results {
		str += fmt.Sprintf("📌 %s\n\n", result.Title)
		//str += fmt.Sprintf("%s\n\n", result.Description)
		//str += fmt.Sprintf("%s\n", result.Content)
		str += fmt.Sprintf("📎 %s\n\n", result.URL)
		str += fmt.Sprintf("\n")
	}

	return str
}
