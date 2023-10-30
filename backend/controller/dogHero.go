package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	service "github.com/yuyuancha/yuyuancha-tool/service/dogHero"
)

// DogHero dog hero controller
type DogHero struct{}

// GetMonthlyTargetList 取得月度達標清單
func (ctrl *DogHero) GetMonthlyTargetList(context *gin.Context) {
	var logic = new(service.DogHeroMonthlyTargetLogic)
	var list, err = logic.GetTargetList()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code":    900001,
			"message": fmt.Sprintf("取得月度達標 .csv 失敗: %s", err.Error()),
			"data":    nil,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "Success",
		"data":    list,
	})
}
