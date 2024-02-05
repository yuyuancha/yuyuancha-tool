package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yuyuancha/yuyuancha-tool/config"
	"github.com/yuyuancha/yuyuancha-tool/service"
	"net/http"
	"strconv"
)

var DebugController Common

// Debug debug Controller
type Debug struct{}

// SetGovTravelCardLocationByShopId 透過店家 id 設定政府旅遊卡位置
func (ctrl *Debug) SetGovTravelCardLocationByShopId(ctx *gin.Context) {
	debugLogic := new(service.DebugLogic)

	shopId := ctx.Param("id")
	id, err := strconv.Atoi(shopId)
	if err != nil || id < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    config.ErrorCodeFormatValid,
			"message": "店家ID" + config.ErrorCodeMessage[config.ErrorCodeFormatValid],
		})
		return
	}

	err = debugLogic.SetGovTravelCardLocationLatAndLonById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": config.ErrorCodeBadRequest,
			"message": fmt.Sprintf(
				"%s: %s",
				config.ErrorCodeMessage[config.ErrorCodeBadRequest],
				err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "成功",
	})
}

// SetUnsettedGovTravelCardLocation 設定政府旅遊卡位置
func (ctrl *Debug) SetUnsettedGovTravelCardLocation(ctx *gin.Context) {
	debugLogic := new(service.DebugLogic)

	err := debugLogic.SetUnsettedGovTravelCardLocation()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": config.ErrorCodeBadRequest,
			"message": fmt.Sprintf(
				"%s: %s",
				config.ErrorCodeMessage[config.ErrorCodeBadRequest],
				err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "成功",
	})
}

// ScrapyGovTravelCardShops 爬取政府旅遊卡店家資料
func (ctrl *Debug) ScrapyGovTravelCardShops(ctx *gin.Context) {
	debugLogic := new(service.DebugLogic)

	from, _ := ctx.GetQuery("from")
	to, _ := ctx.GetQuery("to")

	err := debugLogic.ScrapyGovTravelCardShop(from, to)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": config.ErrorCodeBadRequest,
			"message": fmt.Sprintf(
				"%s: %s",
				config.ErrorCodeMessage[config.ErrorCodeBadRequest],
				err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "成功",
	})
}
