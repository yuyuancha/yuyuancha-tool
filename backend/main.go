package main

import (
	"github.com/yuyuancha/yuyuancha-tool/server"
	"log"
	"time"
)

func init() {
	location, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		log.Fatalln("設定時區發生錯誤:", err.Error())
	}
	time.Local = location
}

func main() {
	server.Run()
}
