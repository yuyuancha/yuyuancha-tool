package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	service "github.com/yuyuancha/yuyuancha-tool/service/weather"
)

// Weather 天氣 controller
type Weather struct{}

// GetOneWeek 取得一週天氣資料
func (ctrl *Weather) GetOneWeek(context *gin.Context) {
	var logic = new(service.WeatherLogic)
	var locations, err = logic.GetOneWeek()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code":    900001,
			"message": fmt.Sprintf("一週天氣資料失敗: %s", err.Error()),
			"data":    nil,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "Success",
		"data":    locations,
	})
}
