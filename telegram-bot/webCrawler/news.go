package webCrawler

import (
	"context"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/yuyuancha/yuyuancha-tool/telegram-bot/model"
	"strconv"
	"strings"
)

const (
	newsURL = "https://news.google.com/topics/CAAqKggKIiRDQkFTRlFvSUwyMHZNRFZxYUdjU0JYcG9MVlJYR2dKVVZ5Z0FQAQ?hl=zh-TW&gl=TW&ceid=TW:zh-Hant"
)

// GetDailyTopHeadlines 取得每日頭條
func GetDailyTopHeadlines() []*model.NewsArticles {
	ctx := context.Background()
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", false),
		chromedp.Flag("hide-scrollbars", false),
		chromedp.Flag("mute-audio", false),
	}

	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

	execAllocatorCtx, execAllocatorCancel := chromedp.NewExecAllocator(ctx, options...)
	defer execAllocatorCancel()

	ctx, cancel := chromedp.NewContext(execAllocatorCtx)
	defer cancel()

	return getNewsByGoogleNews(ctx)
}

// 透過 Google News 取得新聞
func getNewsByGoogleNews(ctx context.Context) []*model.NewsArticles {
	var nodes []*cdp.Node
	var links []string
	var articles []*model.NewsArticles
	_ = chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(newsURL),
		chromedp.WaitReady(`main > c-wiz > c-wiz > c-wiz:nth-child(1) > c-wiz > div > article > a`),
		chromedp.Nodes(`main > c-wiz > c-wiz > c-wiz > c-wiz > div > article > div.XlKvRb > a`, &nodes),
	})

	for _, node := range nodes {
		link := strings.Replace(node.AttributeValue("href"), ".", "https://news.google.com", 1)
		links = append(links, link)
	}

	for i := 1; i <= 5; i++ {
		if len(links) < i {
			break
		}

		var title string
		_ = chromedp.Run(ctx, chromedp.Tasks{
			chromedp.Text(`main > c-wiz > c-wiz > c-wiz:nth-child(`+strconv.Itoa(i)+`) > c-wiz > div > article > a`, &title),
		})

		articles = append(articles, &model.NewsArticles{
			Title: title,
			URL:   links[i-1],
		})
	}

	return articles
}
