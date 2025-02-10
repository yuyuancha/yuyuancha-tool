package service

import (
	"errors"
	"github.com/yuyuancha/yuyuancha-tool/telegram-bot/model"
	"log"
)

// 建立會員
func createUser(user *model.User) error {
	chatId := user.ChatId

	if user.IsExist(chatId) {
		return errors.New("您已經是會員了。")
	}

	err := user.Create()
	if err != nil {
		log.Println("註冊失敗:", err)
		return errors.New("註冊失敗。")
	}

	return nil
}
