package apiCaller

import "github.com/yuyuancha/yuyuancha-tool/telegram-bot/model"

// NewsApiDailyTopHeadlinesResponse 每日頭條回傳格式
type NewsApiDailyTopHeadlinesResponse struct {
	Status       string           `json:"status"`
	TotalResults int              `json:"totalResults"`
	Articles     []NewsApiArticle `json:"articles"`
}

// ToModel 轉換為模型
func (response *NewsApiDailyTopHeadlinesResponse) ToModel() []*model.NewsArticles {
	var results []*model.NewsArticles
	for _, article := range response.Articles {
		articlePointer := &article
		results = append(results, articlePointer.ToModel())
	}
	return results
}

// NewsApiArticle 新聞文章
type NewsApiArticle struct {
	Source      NewsApiSource `json:"source"`
	Author      string        `json:"author"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	URL         string        `json:"url"`
	URLToImage  string        `json:"urlToImage"`
	Content     string        `json:"content"`
}

// ToModel 轉換為模型
func (article *NewsApiArticle) ToModel() *model.NewsArticles {
	return &model.NewsArticles{
		Source:      article.Source.Name,
		Author:      article.Author,
		Title:       article.Title,
		Description: article.Description,
		URL:         article.URL,
		URLToImage:  article.URLToImage,
		Content:     article.Content,
	}
}

// NewsApiSource 新聞來源
type NewsApiSource struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
