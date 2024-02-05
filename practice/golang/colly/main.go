package main

import (
	"fmt"
	"github.com/gocolly/colly/v2/extensions"
	"time"

	"github.com/gocolly/colly/v2"
)

var baseURL = "https://travel.nccc.com.tw/NASApp/NTC/servlet/com.du.mvc.EntryServlet?Action=RetailerList"

func main() {
	url := baseURL

	c := colly.NewCollector()

	extensions.RandomUserAgent(c)
	c.DetectCharset = true
	//c.Async = true

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 2 * time.Second,
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("Visited: %s\n", r.Request.URL)
	})

	c.OnHTML("tr[bgcolor='#FAFAF5']", func(e *colly.HTMLElement) {
		strs := e.ChildTexts("td small")
		//fmt.Println("e:", e.Text)
		fmt.Println("strs:", strs)
	})

	// Setting up an error callback
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println(fmt.Sprintf("Error: %s: Request URL: %s", err, r.Request.URL))
	})

	// Visiting the specified URL to fetch holidays data
	err := c.Post(url, map[string]string{
		"Type": "GetFull",
	})
	//err := c.Visit(url)
	if err != nil {
		fmt.Println("Error visiting URL:", err.Error())
	}

	c.Wait()
}
