package service

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/yuyuancha/yuyuancha-tool/telegram-bot/model"
	"os"
	"strings"
)

// 測試
func (service *Telegram) test(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "測試成功")
	msg.ReplyMarkup = tgbotapi.ReplyKeyboardMarkup{
		Keyboard: [][]tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("股票 0050"),
				tgbotapi.NewKeyboardButton("股票 台積電"),
				tgbotapi.NewKeyboardButton("股票 鴻海"),
			),
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("功能"),
				tgbotapi.NewKeyboardButton("機票"),
			),
		},
	}
	service.send(msg)
}

// 取得新聞
func (service *Telegram) getNews(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "快訊／台大醫院公費流感疫苗打光了！即起暫停接種服務")
	service.send(msg)
}

// 打招呼
func (service *Telegram) greet(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID,
		fmt.Sprintf("你好，%s！", update.Message.Chat.FirstName))
	service.send(msg)
}

// 註冊
func (service *Telegram) register(update tgbotapi.Update) {
	var (
		chatId = update.Message.Chat.ID
		user   = model.User{
			ChatId:    int(chatId),
			FirstName: update.Message.Chat.FirstName,
			Username:  update.Message.Chat.UserName,
			Status:    model.UserStatusActive,
			Role:      model.UserRoleUser,
		}
		msgText = "註冊成功。"
	)

	if err := createUser(&user); err != nil {
		msgText = err.Error()
	}

	msg := tgbotapi.NewMessage(chatId, msgText)
	service.send(msg)
}

// 取得股票
func (service *Telegram) getStocks(update tgbotapi.Update) {
	search := strings.TrimPrefix(update.Message.Text, "股票 ")
	var stock = &model.DailyTradeStock{}
	stock = stock.FindByCodeOrName(search)
	if stock.Id == 0 {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "查無此股票")
		service.send(msg)
		return
	}
	changeStr := fmt.Sprintf("📈⬆️%.2f", stock.Change)
	if stock.Change < 0 {
		changeStr = fmt.Sprintf("📉⬇️%.2f", stock.Change*-1)
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf(
		"股票: %s %s\n開盤價: %.2f\n最高價: %.2f\n最低價: %.2f\n收盤價: %.2f\n成交量: %d\n漲跌: %s",
		stock.Code, stock.Name, stock.OpeningPrice, stock.HighestPrice, stock.LowestPrice, stock.ClosingPrice, stock.TradeVolume, changeStr))
	service.send(msg)
}

// 取得股票
func (service *Telegram) getStocksQuery(update tgbotapi.Update) {
	search := strings.TrimPrefix(update.CallbackQuery.Data, "股票 ")
	var stock = &model.DailyTradeStock{}
	stock = stock.FindByCodeOrName(search)
	if stock.Id == 0 {
		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "查無此股票")
		service.send(msg)
		return
	}
	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, fmt.Sprintf(
		"股票: %s %s\n開盤價: %.2f\n最高價: %.2f\n最低價: %.2f\n收盤價: %.2f\n成交量: %d",
		stock.Code, stock.Name, stock.OpeningPrice, stock.HighestPrice, stock.LowestPrice, stock.ClosingPrice, stock.TradeVolume))
	service.send(msg)
}

// 取得機票
func (service *Telegram) getTicket(update tgbotapi.Update) {
	img, err := os.ReadFile("assets/image/ticket.png")
	if err != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "檔案讀取失敗。")
		service.send(msg)
		return
	}

	photo := tgbotapi.FileBytes{Name: "ticket.png", Bytes: img}
	msg := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, photo)
	service.send(msg)
}

// 取得功能
func (service *Telegram) getFunctions(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf(
		"🔈你好，%s！請選擇您想要執行的功能快捷按鈕😇！", update.Message.Chat.FirstName))
	msg.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("📈台積電股票", "股票 台積電"),
				tgbotapi.NewInlineKeyboardButtonData("📈台灣50股票", "股票 0050"),
				tgbotapi.NewInlineKeyboardButtonData("📈鴻海股票", "股票 鴻海"),
			),
		},
	}
	service.send(msg)
}

// 詢問 AI 服務
func (service *Telegram) askQuestionToAI(update tgbotapi.Update) {
	search := strings.TrimPrefix(update.Message.Text, "ai ")
	result := requestAI(search)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, result)
	service.send(msg)
}
