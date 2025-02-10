package driver

import (
	"fmt"
	"github.com/yuyuancha/yuyuancha-tool/telegram-bot/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MySql *gorm.DB

func init() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Driver.Mysql.User,
		config.Driver.Mysql.Password,
		config.Driver.Mysql.Host,
		config.Driver.Mysql.Port,
		config.Driver.Mysql.Database)

	MySql, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("連線 MySql 資料庫失敗:", err.Error())
	}
}
