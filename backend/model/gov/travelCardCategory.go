package model

import (
	"github.com/yuyuancha/yuyuancha-tool/driver"
	"time"
)

// GovTravelCardCategory 政府旅遊卡類別
type GovTravelCardCategory struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	UpdateTime time.Time `json:"update_time"`
	CreateTime time.Time `json:"create_time"`
}

func (model *GovTravelCardCategory) TableName() string {
	return "gov_travel_card_categories"
}

// FindAll 查詢所有政府旅遊卡類別
func (model *GovTravelCardCategory) FindAll() ([]GovTravelCardCategory, error) {
	var categories []GovTravelCardCategory
	err := driver.MySql.Find(&categories).Error

	return categories, err
}
