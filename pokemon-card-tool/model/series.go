package model

import (
	"github.com/yuyuancha-tool/pokemon-card-tool/driver"
	"time"
)

// Series 系列
type Series struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Code       string    `json:"code"`
	BaseCount  int       `json:"base_count"`
	ExtraCount int       `json:"extra_count"`
	UpdateTime time.Time `json:"update_time"`
	CreateTime time.Time `json:"create_time"`
}

// TableName 回傳 table 名稱
func (s *Series) TableName() string {
	return "series"
}

// GetBaseCountById 透過 ID 取得基礎卡數
func (s *Series) GetBaseCountById(id int) (int, error) {
	err := driver.MySql.Table(s.TableName()).
		Where("id = ?", id).
		Select("base_count").
		First(s).Error
	if err != nil {
		return 0, err
	}

	return s.BaseCount, nil
}
