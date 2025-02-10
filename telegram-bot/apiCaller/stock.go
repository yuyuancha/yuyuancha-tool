package apiCaller

import (
	"encoding/json"
	"github.com/yuyuancha/yuyuancha-tool/telegram-bot/model"
	"io"
	"net/http"
	"strconv"
)

// Stock 股票服務
var Stock = &StockCaller{}

// StockCaller 請求股票服務
type StockCaller struct{}

var stockUrl = "https://openapi.twse.com.tw/v1/exchangeReport/STOCK_DAY_ALL"

// UpdateDailyTradeStocks 更新每日交易股票
func (caller *StockCaller) UpdateDailyTradeStocks() {
	response, err := http.Get(stockUrl)
	if err != nil {
		panic(err)
	}

	defer func() {
		err = response.Body.Close()
		if err != nil {
			panic(err)
		}
	}()

	if response.StatusCode != http.StatusOK {
		panic("status code error: " + response.Status)
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var results []DailyTradeStockResponse

	err = json.Unmarshal(bodyBytes, &results)
	if err != nil {
		panic(err)
	}

	for _, result := range results {
		var stock = &model.DailyTradeStock{}
		stock.Code = result.Code
		stock.Name = result.Name
		stock.TradeVolume, _ = strconv.Atoi(result.TradeVolume)
		stock.TradeValue, _ = strconv.Atoi(result.TradeValue)
		stock.OpeningPrice, _ = strconv.ParseFloat(result.OpeningPrice, 64)
		stock.HighestPrice, _ = strconv.ParseFloat(result.HighestPrice, 64)
		stock.LowestPrice, _ = strconv.ParseFloat(result.LowestPrice, 64)
		stock.ClosingPrice, _ = strconv.ParseFloat(result.ClosingPrice, 64)
		stock.Change, _ = strconv.ParseFloat(result.Change, 64)
		stock.Transaction, _ = strconv.Atoi(result.Transaction)
		err = stock.Create()
		if err != nil {
			panic(err)
		}
	}
}
