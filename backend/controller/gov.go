package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yuyuancha/yuyuancha-tool/config"
	service "github.com/yuyuancha/yuyuancha-tool/service/gov"
	"net/http"
)

// Gov 政府相關 controller
type Gov struct{}

// GetTravelCardCategories 取得政府旅遊卡類別
func (ctrl *Gov) GetTravelCardCategories(context *gin.Context) {
	var logic = new(service.GovLogic)
	var list, err = logic.GetTravelCategories()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    config.ErrorCodeBadRequest,
			"message": config.ErrorCodeMessage[config.ErrorCodeBadRequest],
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "成功",
		"data":    list,
	})
}

// GetTravelCardShops 取得政府旅遊卡地點清單
func (ctrl *Gov) GetTravelCardShops(context *gin.Context) {
	var form = struct {
		CategoryId int `json:"categoryId"`
	}{}
	err := context.ShouldBindJSON(&form)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    config.ErrorCodeBadRequest,
			"message": config.ErrorCodeMessage[config.ErrorCodeBadRequest],
		})
		return
	}

	var logic = new(service.GovLogic)
	list, err := logic.GetGovTravelCardShops(form.CategoryId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    config.ErrorCodeBadRequest,
			"message": config.ErrorCodeMessage[config.ErrorCodeBadRequest],
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "成功",
		"data": gin.H{
			"count": len(list),
			"list":  list,
		},
	})
}
