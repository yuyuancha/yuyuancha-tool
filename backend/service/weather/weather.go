package weather

import (
	"github.com/yuyuancha/yuyuancha-tool/apiCaller"
	model "github.com/yuyuancha/yuyuancha-tool/model/weather"
)

// WeatherLogic 天氣邏輯
type WeatherLogic struct{}

var weatherApiCaller apiCaller.WeatherApiCaller

func init() {
	weatherApiCaller.Init()
}

// GetOneWeek 取得一週天氣資料
func (logic *WeatherLogic) GetOneWeek() ([]model.WeatherOneWeekLocation, error) {
	return weatherApiCaller.GetOneWeek()
}
