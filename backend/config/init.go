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

// appConfigStruct 程序環境 Config 結構
type appConfigStruct struct {
	Environment string
	Port        int
}

// mysqlStruct MySql Config 結構
type mysqlStruct struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
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
		Port:        v.GetInt("APP_PORT"),
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
}
