package model

import (
	"github.com/yuyuancha/yuyuancha-tool/telegram-bot/driver"
	"time"
)

// DailyTradeStock 日成交股票
type DailyTradeStock struct {
	Id           int       `json:"id" gorm:"pk"`
	Code         string    `json:"code"`
	Name         string    `json:"name"`
	TradeVolume  int       `json:"tradeVolume"`
	TradeValue   int       `json:"tradeValue"`
	OpeningPrice float64   `json:"openingPrice"`
	HighestPrice float64   `json:"highestPrice"`
	LowestPrice  float64   `json:"lowestPrice"`
	ClosingPrice float64   `json:"closingPrice"`
	Change       float64   `json:"change"`
	Transaction  int       `json:"transaction"`
	UpdateTime   time.Time `json:"updateTime" gorm:"autoUpdateTime"`
	CreateTime   time.Time `json:"createTime" gorm:"autoCreateTime"`
}

// TableName 資料表名稱
func (stock *DailyTradeStock) TableName() string {
	return "daily_trade_stocks"
}

// Create 建立紀錄
func (stock *DailyTradeStock) Create() error {
	var s *DailyTradeStock
	driver.MySql.Where("code = ?", stock.Code).First(&s)
	if s.Id > 0 {
		return driver.MySql.Model(&s).Updates(stock).Error
	}
	return driver.MySql.Create(stock).Error
}

// FindByCodeOrName 透過股票代碼或名稱尋找股票
func (stock *DailyTradeStock) FindByCodeOrName(search string) *DailyTradeStock {
	driver.MySql.Where("code = ? or name = ?", search, search).First(&stock)
	return stock
}
