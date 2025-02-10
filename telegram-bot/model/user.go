package model

import (
	"github.com/yuyuancha/yuyuancha-tool/telegram-bot/driver"
	"time"
)

const (
	UserStatusBlock  = 0
	UserStatusActive = 1

	UserRoleUser  = 1
	UserRoleAdmin = 2
)

// User 會員
type User struct {
	Id         int       `json:"id" gorm:"pk"`
	ChatId     int       `json:"chatId"`
	FirstName  string    `json:"firstName"`
	Username   string    `json:"username"`
	Status     int       `json:"status"`
	Role       int       `json:"role"`
	UpdateTime time.Time `json:"updateTime" gorm:"autoUpdateTime"`
	CreateTime time.Time `json:"createTime" gorm:"autoCreateTime"`
}

// TableName 資料表名稱
func (user *User) TableName() string {
	return "users"
}

// IsExist 判斷會員是否存在
func (user *User) IsExist(chatId int) bool {
	driver.MySql.Where("chat_id = ?", chatId).First(&user)
	return user.Id > 0
}

// IsAdmin 判斷是否為管理員
func (user *User) IsAdmin(chatId int) bool {
	driver.MySql.Where("chat_id = ? AND role = ?", chatId, UserRoleAdmin).First(&user)
	return user.Id > 0
}

// IsBlock 判斷是否被封鎖
func (user *User) IsBlock(chatId int) bool {
	var u *User
	driver.MySql.Where("chat_id = ? AND status = ?", chatId, UserStatusBlock).First(&u)
	return u.Id > 0
}

// AllActive 取得所有有效會員
func (user *User) AllActive() []*User {
	var users []*User
	driver.MySql.Where("status = ?", UserStatusActive).Find(&users)
	return users
}

// FindByChatId 透過 chatId 尋找會員
func (user *User) FindByChatId(chatId int) *User {
	driver.MySql.Where("chat_id = ?", chatId).First(&user)
	return user
}

// Create 建立紀錄
func (user *User) Create() error {
	return driver.MySql.Create(user).Error
}
