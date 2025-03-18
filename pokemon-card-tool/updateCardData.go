package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/yuyuancha-tool/pokemon-card-tool/contants"
	"github.com/yuyuancha-tool/pokemon-card-tool/webCrawler"
	"slices"
)

func main() {
	var seriesId int
	flag.IntVar(&seriesId, "series", contants.SeriesIdA2, "輸入系列 ID: 1. A2系列-時間激鬥(預設) 2. A2a系列-超克之光")
	flag.Parse()
	err := validateSeriesId(seriesId)
	if err != nil {
		panic(err)
	}

	c := webCrawler.NewPokemonCardCrawler()
	err = c.CreateCardsDataBySeries(contants.SeriesIdA2a)
	if err != nil {
		panic(err)
	}
	fmt.Println("done")
}

// 驗證系列 ID
func validateSeriesId(seriesId int) error {
	if slices.Index([]int{contants.SeriesIdA2, contants.SeriesIdA2a}, seriesId) == -1 {
		return errors.New("無效的系列 ID")
	}
	return nil
}
