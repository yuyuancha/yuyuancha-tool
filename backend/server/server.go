package server

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yuyuancha/yuyuancha-tool/config"
	"log"
)

// Run 啟動伺服器服務
func Run() {
	router := gin.Default()
	router.Use(cors.Default())

	setRouter(router)

	err := router.Run(fmt.Sprintf(":%s", config.App.Port))
	if err != nil {
		log.Fatalln("開啟 Gin 服務失敗:", err.Error())
	}
}
