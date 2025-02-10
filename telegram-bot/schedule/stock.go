package schedule

import (
	"github.com/go-co-op/gocron"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/yuyuancha/yuyuancha-tool/telegram-bot/service"
	"time"
)

// RunSchedule 啟動排程
func RunSchedule() {
	scheduler := gocron.NewScheduler(time.Local)

	_, err := scheduler.Every(1).Day().At("14:00").Do(func() {
		service.UpdateDailyTradeStocks()
		service.TelegramService.SendAll(tgbotapi.NewMessage(0, "⏱️✅成功更新股票排程！"))
	})
	if err != nil {
		service.TelegramService.SendAll(tgbotapi.NewMessage(0, "新增股票排程失敗: "+err.Error()))
	}

	go func() {
		scheduler.StartAsync()
	}()
}
