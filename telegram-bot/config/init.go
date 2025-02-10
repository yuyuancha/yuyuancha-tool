package config

import (
	"log"

	"github.com/spf13/viper"
)

// App 主要程序環境 Config
var App appConfigStruct

// Driver 資料庫環境 Config
var Driver struct {
	Mysql mysqlStruct
}

// TelegramBot telegram 機器人結構
var TelegramBot telegramBotStruct

// GoogleService google service 結構
var GoogleService googleServiceStruct

// 程序環境 Config 結構
type appConfigStruct struct {
	Environment string
}

// MySql Config 結構
type mysqlStruct struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

// telegram 機器人結構
type telegramBotStruct struct {
	Token string
}

// google service 結構
type googleServiceStruct struct {
	GminiApiKey string
}

func init() {
	v := viper.New()
	v.SetConfigName(".env")
	v.SetConfigType("env")
	v.AddConfigPath(".")

	err := v.ReadInConfig()
	if err != nil {
		log.Fatalln("無法讀取配置文件:", err.Error())
	}

	// 配置主要程序環境 Config
	App = appConfigStruct{
		Environment: v.GetString("APP_ENV"),
	}

	// 配置資料庫環境 Config
	Driver = struct{ Mysql mysqlStruct }{
		Mysql: mysqlStruct{
			Host:     v.GetString("MYSQL_HOST"),
			Port:     v.GetInt("MYSQL_PORT"),
			User:     v.GetString("MYSQL_USER"),
			Password: v.GetString("MYSQL_PASSWORD"),
			Database: v.GetString("MYSQL_DATABASE"),
		},
	}

	// 配置 telegram 機器人結構
	TelegramBot = telegramBotStruct{
		Token: v.GetString("TG_BOT_TOKEN"),
	}

	// 配置 google service 結構
	GoogleService = googleServiceStruct{
		GminiApiKey: v.GetString("GOOGLE_GMINI_API_KEY"),
	}
}
