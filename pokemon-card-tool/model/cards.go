package model

import (
	"fmt"
	"github.com/yuyuancha-tool/pokemon-card-tool/driver"
	"strings"
	"time"
)

// Card 寶可夢卡牌
type Card struct {
	ID         int        `json:"id"`
	Number     string     `json:"number"`
	SeriesId   int        `json:"series_id"`
	Name       string     `json:"name"`
	Attribute  string     `json:"attribute"`
	Rarity     string     `json:"rarity"`
	CreateTime time.Time  `json:"create_time" gorm:"autoCreateTime"`
	UpdateTime time.Time  `json:"update_time" gorm:"autoUpdateTime"`
	Packages   []*Package `json:"packages" gorm:"many2many:package_card_relations;"`
}

// TableName 回傳 table 名稱
func (c *Card) TableName() string {
	return "cards"
}

// CreateCards 新增卡牌
func (c *Card) CreateCards(cards []*Card) error {
	tx := driver.MySql.Begin()

	for _, card := range cards {
		err := tx.Create(card).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

// GetCardsByNumbers 透過編號取得卡牌
func (c *Card) GetCardsByNumbers(seriesId int, numbers []string) ([]*Card, error) {
	var cards []*Card
	err := driver.MySql.Table(c.TableName()).
		Where("series_id = ?", seriesId).
		Where("number in (?)", numbers).
		Preload("Packages").
		Find(&cards).Error
	if err != nil {
		return nil, err
	}

	return cards, nil
}

// GetCardString 取得卡片字串
func (c *Card) GetCardString() string {
	return fmt.Sprintf("卡片編號：%s, 卡片名稱：%s, 卡片屬性：%s, 卡片稀有度：%s, 擴充包：%s", c.Number, c.Name, c.Attribute, c.Rarity, c.GetPackagesString())
}

// GetPackagesString 取得擴充包字串
func (c *Card) GetPackagesString() string {
	if len(c.Packages) == 0 {
		return "尚未存在擴充包資料"
	}
	var packageNames []string
	for _, p := range c.Packages {
		packageNames = append(packageNames, p.Name)
	}
	return strings.Join(packageNames, ", ")
}
