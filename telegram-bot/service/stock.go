package service

import (
	"github.com/yuyuancha/yuyuancha-tool/telegram-bot/apiCaller"
)

// UpdateDailyTradeStocks 更新每日交易股票
func UpdateDailyTradeStocks() {
	apiCaller.Stock.UpdateDailyTradeStocks()
}
