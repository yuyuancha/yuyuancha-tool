package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://travel.nccc.com.tw/NASApp/NTC/servlet/com.du.mvc.EntryServlet?Action=RetailerInquiry&Type=GetFull", nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// 提取文章標題
	doc.Find("tr[bgcolor='#EEEEEE'] td small").Each(func(index int, item *goquery.Selection) {
		title := item.Text()
		fmt.Println(title)
	})

	fmt.Println("done!")
}
