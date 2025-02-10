package apiCaller

// DailyTradeStockResponse 每日交易股票回傳格式
type DailyTradeStockResponse struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	TradeVolume  string `json:"tradeVolume"`
	TradeValue   string `json:"tradeValue"`
	OpeningPrice string `json:"openingPrice"`
	HighestPrice string `json:"highestPrice"`
	LowestPrice  string `json:"lowestPrice"`
	ClosingPrice string `json:"closingPrice"`
	Change       string `json:"change"`
	Transaction  string `json:"transaction"`
}
