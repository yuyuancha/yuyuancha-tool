package main

import (
	"github.com/yuyuancha/yuyuancha-tool/telegram-bot/schedule"
	"github.com/yuyuancha/yuyuancha-tool/telegram-bot/service"
	"time"
)

func init() {
	time.Local, _ = time.LoadLocation("Asia/Taipei")
}

func main() {
	schedule.RunSchedule()
	service.RunTelegramService()
}
