package webCrawler

import (
	"context"
	"errors"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/yuyuancha-tool/pokemon-card-tool/contants"
	"github.com/yuyuancha-tool/pokemon-card-tool/model"
	"strings"
	"time"
)

// PokemonCardCrawler PokemonCard 結構
type PokemonCardCrawler struct{}

// NewPokemonCardCrawler new PokemonCard
func NewPokemonCardCrawler() *PokemonCardCrawler {
	return &PokemonCardCrawler{}
}

// CreateCardsDataBySeries 建立系列卡片資料
func (crawler *PokemonCardCrawler) CreateCardsDataBySeries(seriesId int) error {
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

	timeoutCtx, timeoutCancel := context.WithTimeout(ctx, 10*time.Second)
	defer timeoutCancel()

	var (
		tableChildId   = contants.SeriesWebCrawlerChildrenId[seriesId]
		tbodySelPrefix = fmt.Sprintf(`#mw-content-text > div.mw-content-ltr.mw-parser-output > table:nth-child(%d) > tbody > tr:nth-child(2) > td > table > tbody`, tableChildId)
		tbody          []*cdp.Node
		count          int64
	)

	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(contants.SeriesWebCrawlerUrl[seriesId]),
		chromedp.WaitReady(tbodySelPrefix),
		chromedp.Nodes(tbodySelPrefix, &tbody),
	})
	if err != nil {
		return errors.New(fmt.Sprintf("爬蟲抓取資料錯誤：%s", err.Error()))
	}

	if len(tbody) == 0 {
		return errors.New("查無 tbody 結構")
	}

	count = tbody[0].ChildNodeCount
	fmt.Println("抓取到的子結構數量：", count)

	var cards []*model.Card
	var packageModel model.Package

	var seriesModel model.Series
	baseCount, err := seriesModel.GetBaseCountById(seriesId)
	if err != nil {
		return errors.New(fmt.Sprintf("取得基礎卡數錯誤：%s", err.Error()))
	}

	var packages = packageModel.GetPackagesBySeriesId(seriesId)
	if len(packages) == 0 {
		return errors.New("資料庫查無擴充包")
	}

	for i := 2; i <= int(count); i++ {
		if i > baseCount+1 {
			break
		}
		var (
			selPrefix      = fmt.Sprintf(`%s > tr:nth-child(%d) > `, tbodySelPrefix, i)
			number         string
			nameLinks      []*cdp.Node
			attributeNodes []*cdp.Node
			rankNodes      []*cdp.Node
			packageNodes   []*cdp.Node
		)

		err = chromedp.Run(ctx, chromedp.Tasks{
			chromedp.Text(fmt.Sprintf(`%std:nth-child(1)`, selPrefix), &number),
			chromedp.Nodes(fmt.Sprintf(`%std:nth-child(2) > a`, selPrefix), &nameLinks),
			chromedp.Nodes(fmt.Sprintf(`%std:nth-child(4)`, selPrefix), &rankNodes),
		})
		if err != nil {
			return errors.New(fmt.Sprintf("爬蟲抓取資料(index: %d)錯誤：%s", i, err.Error()))
		}

		if seriesId != contants.SeriesIdA2a {
			err = chromedp.Run(ctx, chromedp.Tasks{
				chromedp.Nodes(fmt.Sprintf(`%std:nth-child(5) > div > span > span > span`, selPrefix), &packageNodes),
			})
			if err != nil {
				return errors.New(fmt.Sprintf("爬蟲抓取資料(index: %d)擴充包錯誤：%s", i, err.Error()))
			}
		}

		// 道具、物品的屬性欄位沒有 span tag
		err = chromedp.Run(timeoutCtx, chromedp.ActionFunc(func(ctx context.Context) error {
			err = chromedp.Nodes(fmt.Sprintf(`%sth > span > a`, selPrefix), &attributeNodes).Do(ctx)

			return err
		}))

		if err != nil {
			err = chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
				err = chromedp.Nodes(fmt.Sprintf(`%sth > a`, selPrefix), &attributeNodes).Do(ctx)

				return err
			}))
			if err != nil {
				return errors.New(fmt.Sprintf("爬蟲抓取資料(index: %d)屬性錯誤：%s", i, err.Error()))
			}
		}

		if len(nameLinks) == 0 {
			continue
		}
		name, _ := nameLinks[0].Attribute("title")

		if len(attributeNodes) == 0 {
			return errors.New(fmt.Sprintf("爬蟲抓取資料(index: %d)錯誤：查無屬性", i))
		}
		attribute, _ := attributeNodes[0].Attribute("title")

		if len(rankNodes) == 0 {
			return errors.New(fmt.Sprintf("爬蟲抓取資料(index: %d)錯誤：查無稀有度節點", i))
		}
		if rankNodes[0].ChildNodeCount == 0 {
			return errors.New(fmt.Sprintf("爬蟲抓取資料(index: %d)錯誤：查無稀有度子節點數量", i))
		}
		rankCount := rankNodes[0].ChildNodeCount

		var packagesName []string
		for _, packageNode := range packageNodes {
			packageTitle, _ := packageNode.Attribute("title")
			packagesName = append(packagesName, packageTitle)
		}

		// 將編號尾綴去除 Ex: 001/021 -> 001
		SplitNumbers := strings.Split(number, "/")

		// 將名字尾綴去除 Ex: 走路草（A2）-> 走路草
		splitName := strings.Split(name, "（")

		// 將屬性尾綴去除 Ex: 草（TCG）-> 草
		splitAttribute := strings.Split(attribute, "（")

		cards = append(cards, &model.Card{
			Number:     SplitNumbers[0],
			Name:       splitName[0],
			SeriesId:   seriesId,
			Attribute:  splitAttribute[0],
			Rarity:     crawler.getPtByRank(int(rankCount)),
			Packages:   crawler.comparePackages(packages, packagesName),
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		})
	}

	var cardModel *model.Card
	err = cardModel.CreateCards(cards)
	if err != nil {
		return errors.New(fmt.Sprintf("新增卡片錯誤：%s", err.Error()))
	}

	return nil
}

// 透過稀有度取得點數
func (crawler *PokemonCardCrawler) getPtByRank(rank int) string {
	switch rank {
	case 1:
		return "35pt"
	case 2:
		return "70pt"
	case 3:
		return "150pt"
	case 4:
		return "500pt"
	default:
		return "未知"
	}
}

// 比對卡片的卡包
func (crawler *PokemonCardCrawler) comparePackages(packages []*model.Package, cardPackages []string) []*model.Package {
	if len(cardPackages) == 0 {
		return packages
	}

	var result []*model.Package
	for _, packageModel := range packages {
		for _, cardPackage := range cardPackages {
			if packageModel.Name == cardPackage {
				result = append(result, packageModel)
			}
		}
	}
	return result
}
