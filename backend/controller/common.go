package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var CommonController Common

// Common 通用 Controller
type Common struct{}

// Ping ping-pong 測試用 Api
func (Common) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "pong",
		"data":    nil,
	})
}
