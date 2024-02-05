package server

import (
	"github.com/gin-gonic/gin"
	"github.com/yuyuancha/yuyuancha-tool/controller"
)

func setRouter(router *gin.Engine) {
	dogHeroController := new(controller.DogHero)
	govController := new(controller.Gov)

	debugController := new(controller.Debug)

	router.GET("/ping", controller.CommonController.Ping)

	v1 := router.Group("/v1")

	dogHero := v1.Group("/dogHero")
	dogHero.POST("/monthlyTarget/list", dogHeroController.GetMonthlyTargetList)

	gov := v1.Group("/gov")
	gov.POST("/travelCard/shops", govController.GetTravelCardShops)
	gov.POST("/travelCard/categories", govController.GetTravelCardCategories)

	debug := v1.Group("/debug")
	debug.GET("/gov/travelCard/scrapy", debugController.ScrapyGovTravelCardShops)
	debug.GET("/gov/travelCard/setLocation", debugController.SetGovTravelCardLocationByShopId)
	debug.GET("/gov/travelCard/setUnsettedLocation", debugController.SetUnsettedGovTravelCardLocation)
}
