package service

import (
	"errors"
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
	"github.com/yuyuancha/yuyuancha-tool/apiCaller"
	"github.com/yuyuancha/yuyuancha-tool/config"
	model "github.com/yuyuancha/yuyuancha-tool/model/gov"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

// DebugLogic debug 邏輯
type DebugLogic struct{}

// SetUnsettedGovTravelCardLocation 設定未設定經緯度店家位置
func (logic *DebugLogic) SetUnsettedGovTravelCardLocation() error {
	var shopModel model.GovTravelCardShop
	shops, err := shopModel.FindAllLocationUnsetted()
	if err != nil {
		return errors.New("查詢未設定經緯度店家資料發生錯誤")
	}

	for _, shop := range shops {
		err = logic.SetGovTravelCardLocationLatAndLonById(shop.Id)
		if err != nil {
			return err
		}

		time.Sleep(500 * time.Millisecond)
	}

	return nil
}

// SetGovTravelCardLocationLatAndLonById 透過店家 id 設定政府旅遊卡位置緯度和經度
func (logic *DebugLogic) SetGovTravelCardLocationLatAndLonById(id int) error {
	var shop model.GovTravelCardShop
	shop.Id = id
	err := shop.FindOneById()
	if err == gorm.ErrRecordNotFound {
		return errors.New("查無此店家資料")
	}
	if err != nil {
		return errors.New(fmt.Sprintf("此店家資料查詢發生錯誤：%s", err.Error()))
	}

	googleMapsApi := apiCaller.NewGoogleMapsApiCaller()
	lat, lng := googleMapsApi.GetAddressLatAndLng(shop.Address)

	shop.Latitude = lat
	shop.Longitude = lng

	return shop.UpdateLatAndLon()
}

// ScrapyGovTravelCardShop 爬取政府旅遊卡店家資料
func (logic *DebugLogic) ScrapyGovTravelCardShop(a, b string) error {
	var startPage = -1
	var endPage = -1

	if aInt, err := strconv.Atoi(a); err == nil {
		startPage = aInt
	}

	if bInt, err := strconv.Atoi(b); err == nil {
		endPage = bInt
	}

	if startPage == -1 || endPage == -1 || startPage > endPage {
		return errors.New("輸入頁碼格式錯誤")
	}

	for i := startPage; i <= endPage; i++ {
		err := logic.scrapyGovTravelCardShopByPage(i)
		if err != nil {
			return err
		}
		time.Sleep(200 * time.Millisecond)
	}

	return nil
}

// scrapyGovTravelCardShopByPage 透過頁數爬取政府旅遊卡店家資料
func (logic *DebugLogic) scrapyGovTravelCardShopByPage(page int) error {
	url := "https://travel.nccc.com.tw/NASApp/NTC/servlet/com.du.mvc.EntryServlet?Action=RetailerList&Type=GetFull&WebMode=&Request=NULL_NULL_NULL_001_NULL_NULL_NULL_NULL_NULL_0_%s_20_37944"
	categoryModel := model.GovTravelCardCategory{}
	categories, err := categoryModel.FindAll()
	if err != nil {
		return errors.New("查詢政府旅遊卡類別發生錯誤")
	}

	c := colly.NewCollector()
	c.DetectCharset = true

	extensions.RandomUserAgent(c)

	err = c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 2 * time.Second,
	})
	if err != nil {
		return errors.New(fmt.Sprintf("設定爬蟲限制發生錯誤：%s", err.Error()))
	}

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("爬蟲瀏覽網頁：", r.Request.URL)
	})

	c.OnHTML("tr[bgcolor='#FAFAF5']", func(e *colly.HTMLElement) {
		strs := e.ChildTexts("td small")
		categoryId := -1

		for _, category := range categories {
			if category.Name == strs[config.GovTravelCardShopTableIndexForCategory] {
				categoryId = category.Id
			}
		}

		if categoryId == -1 {
			fmt.Printf("查無此店家類別：%s \n", strs[config.GovTravelCardShopTableIndexForCategory])
		}

		shop := model.GovTravelCardShop{
			Name:        strs[config.GovTravelCardShopTableIndexForShopName],
			CategoryId:  categoryId,
			Address:     strings.ReplaceAll(strs[config.GovTravelCardShopTableIndexForAddress], " ", ""),
			PhoneNumber: strs[config.GovTravelCardShopTableIndexForPhoneNumber],
			UpdateTime:  time.Now(),
			CreateTime:  time.Now(),
		}

		createErr := shop.Create()
		if createErr != nil {
			fmt.Println("建立店家資料發生錯誤：", err.Error())
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println(fmt.Sprintf("網頁爬蟲發生錯誤: %s 。URL: %s", err, r.Request.URL))
	})

	err = c.Visit(fmt.Sprintf(url, strconv.Itoa(page)))
	if err != nil {
		fmt.Println("Error visiting URL:", err.Error())
	}

	c.Wait()

	return nil
}
