package service

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/yuyuancha/yuyuancha-tool/telegram-bot/config"
	"github.com/yuyuancha/yuyuancha-tool/telegram-bot/model"
	"log"
	"strings"
)

var TelegramService = Telegram{}

// Telegram Telegram 服務
type Telegram struct {
	Bot     *tgbotapi.BotAPI
	Updates tgbotapi.UpdatesChannel
}

// RunTelegramService 執行 Telegram 服務
func RunTelegramService() {
	var err error

	TelegramService.Bot, err = tgbotapi.NewBotAPI(config.TelegramBot.Token)
	if err != nil {
		log.Panic(err)
	}

	TelegramService.Bot.Debug = true

	log.Printf("連接授權機器人(名稱: %s)", TelegramService.Bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	TelegramService.SendAll(tgbotapi.NewMessage(0, "機器人已啟動"))

	TelegramService.Updates, _ = TelegramService.Bot.GetUpdatesChan(u)

	TelegramService.listenUpdates()
}

// SendAll 發送給所有用戶
func (service *Telegram) SendAll(msg tgbotapi.MessageConfig) {
	userModel := model.User{}
	users := userModel.AllActive()
	for _, user := range users {
		msg.ChatID = int64(user.ChatId)
		service.send(msg)
	}
}

// 監聽更新
func (service *Telegram) listenUpdates() {
	for update := range service.Updates {
		if update.Message == nil && update.CallbackQuery == nil {
			continue
		}

		if update.Message != nil {
			if service.handleEventWithoutAuth(update) {
				continue
			}

			service.handleEventWithAuth(update)
		}

		if update.CallbackQuery != nil {
			service.handleCallbackQuery(update)
		}
	}
}

// 發送訊息
func (service *Telegram) send(msg tgbotapi.Chattable) {
	_, err := service.Bot.Send(msg)
	if err != nil {
		log.Println("發送訊息失敗:", err)
	}
}

// 執行不需驗證事件
func (service *Telegram) handleEventWithoutAuth(update tgbotapi.Update) (isFinished bool) {
	isFinished = true

	switch update.Message.Text {
	case "你好":
		service.greet(update)
	case "/register":
		service.register(update)
	case "/test":
		service.test(update)
	default:
		isFinished = false
	}
	return
}

// 執行需驗證事件
func (service *Telegram) handleEventWithAuth(update tgbotapi.Update) {
	var chatId = update.Message.Chat.ID
	if err := service.authUser(int(chatId)); err != nil {
		msg := tgbotapi.NewMessage(chatId, err.Error())
		service.send(msg)
		return
	}

	switch {
	case update.Message.Text == "功能":
		service.getFunctions(update)
	case strings.HasPrefix(update.Message.Text, "股票"):
		service.getStocks(update)
	case strings.HasPrefix(update.Message.Text, "ai"):
		service.askTextQuestionToAI(update)
	case update.Message.Text == "機票":
		service.getTicket(update)
	default:
		service.getFunctions(update)
	}
}

// 處理 callback query
func (service *Telegram) handleCallbackQuery(update tgbotapi.Update) {
	switch {
	case update.CallbackQuery.Data == "功能":
		service.getFunctions(update)
	case strings.HasPrefix(update.CallbackQuery.Data, "股票"):
		service.getStocksQuery(update)
	}
}

// 驗證用戶
func (service *Telegram) authUser(chatId int) error {
	var user = model.User{}

	if !user.IsExist(chatId) {
		return errors.New("請先註冊成為會員。")
	}

	if user.IsBlock(chatId) {
		return errors.New("哎呀，好像發生一些錯誤，請聯繫相關人員。")
	}

	return nil
}
