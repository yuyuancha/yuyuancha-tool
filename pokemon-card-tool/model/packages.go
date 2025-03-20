package model

import (
	"github.com/yuyuancha-tool/pokemon-card-tool/driver"
	"time"
)

// Package 擴充包
type Package struct {
	ID         int       `json:"id"`
	SeriesId   int       `json:"series_id"`
	Name       string    `json:"name"`
	UpdateTime time.Time `json:"update_time"`
	CreateTime time.Time `json:"create_time"`
	Series     Series    `json:"series" gorm:"foreignKey:id"`
}

// TableName 回傳 table 名稱
func (p *Package) TableName() string {
	return "packages"
}

// GetPackagesBySeriesId 依照系列 ID 取得擴充包
func (p *Package) GetPackagesBySeriesId(seriesId int) []*Package {
	var packages []*Package
	driver.MySql.Table(p.TableName()).Where("series_id = ?", seriesId).Find(&packages)
	return packages
}
