package server

import (
	"github.com/gin-gonic/gin"
	"github.com/yuyuancha/yuyuancha-tool/controller"
)

func setRouter(router *gin.Engine) {
	dogHeroController := new(controller.DogHero)
	weatherController := new(controller.Weather)

	router.GET("/ping", controller.CommonController.Ping)

	v1 := router.Group("/v1")

	dogHero := v1.Group("/dogHero")
	dogHero.POST("/monthlyTarget/list", dogHeroController.GetMonthlyTargetList)

	weather := v1.Group("/weather")
	weather.GET("/oneWeek", weatherController.GetOneWeek)
}
