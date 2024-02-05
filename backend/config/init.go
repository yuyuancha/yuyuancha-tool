package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// App 主要程序環境 Config
var App appConfigStruct

// Driver 資料庫環境 Config
var Driver struct {
	Mysql mysqlStruct
}

// appConfigStruct 程序環境 Config 結構
type appConfigStruct struct {
	Environment string
	Port        string
}

// mysqlStruct MySql Config 結構
type mysqlStruct struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

// GoogleMapsApiKey google maps api 金鑰
var GoogleMapsApiKey string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("設定 ENV 發生錯誤:", err.Error())
	}

	// 配置主要程序環境 Config
	App = appConfigStruct{
		Environment: os.Getenv("APP_ENV"),
		Port:        os.Getenv("APP_PORT"),
	}

	// 配置資料庫環境 Config
	Driver = struct{ Mysql mysqlStruct }{
		Mysql: mysqlStruct{
			Host:     os.Getenv("MYSQL_HOST"),
			Port:     os.Getenv("MYSQL_PORT"),
			User:     os.Getenv("MYSQL_USER"),
			Password: os.Getenv("MYSQL_PASSWORD"),
			Database: os.Getenv("MYSQL_DATABASE"),
		},
	}

	// 配置 google maps api 金鑰
	GoogleMapsApiKey = os.Getenv("GOOGLE_MAPS_API_KEY")
}
