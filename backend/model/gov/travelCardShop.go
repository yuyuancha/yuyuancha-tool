package model

import (
	"github.com/yuyuancha/yuyuancha-tool/driver"
	"time"
)

// GovTravelCardShop 政府旅遊卡店家
type GovTravelCardShop struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	PhoneNumber string    `json:"phone_number"`
	CategoryId  int       `json:"category_id"`
	Note        string    `json:"note"`
	UpdateTime  time.Time `json:"update_time"`
	CreateTime  time.Time `json:"create_time"`
}

func (model *GovTravelCardShop) TableName() string {
	return "gov_travel_card_shops"
}

// FindAll 找到所有店家
func (model *GovTravelCardShop) FindAll() ([]GovTravelCardShop, error) {
	var shops []GovTravelCardShop
	err := driver.MySql.Where("latitude != 0 AND longitude != 0").Find(&shops).Error

	return shops, err
}

// FindAllByCategoryId 透過類別 ID 找到店家
func (model *GovTravelCardShop) FindAllByCategoryId(categoryId int) ([]GovTravelCardShop, error) {
	var shops []GovTravelCardShop
	err := driver.MySql.Where("category_id = ? AND latitude != 0 AND longitude != 0", categoryId).Find(&shops).Error

	return shops, err
}

// FindAllLocationUnsetted 找到所有未設定經緯度的店家
func (model *GovTravelCardShop) FindAllLocationUnsetted() ([]GovTravelCardShop, error) {
	var shops []GovTravelCardShop
	err := driver.MySql.Where("latitude = 0 AND longitude = 0").Find(&shops).Error

	return shops, err
}

// FindOneById 透過 ID 找到一筆資料
func (model *GovTravelCardShop) FindOneById() error {
	return driver.MySql.First(model, model.Id).Error
}

// UpdateLatAndLon 更新緯度和經度
func (model *GovTravelCardShop) UpdateLatAndLon() error {
	return driver.MySql.Model(model).Updates(
		GovTravelCardShop{
			Latitude:  model.Latitude,
			Longitude: model.Longitude,
		}).Error
}

// Create 建立一筆資料
func (model *GovTravelCardShop) Create() error {
	return driver.MySql.Create(model).Error
}
