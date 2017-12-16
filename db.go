package main

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

type Impl struct {
	DB *gorm.DB
}

func (i *Impl) InitDB() {
	var err error
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	db := "fake_idol"
	parameter := "charset=utf8&parseTime=True"
	i.DB, err = gorm.Open("mysql", user+":"+pass+"@/"+db+"?"+parameter)
	if err != nil {
		log.Fatalf("DB接続時にエラーが発生しました: '%v'", err)
	}
	i.DB.LogMode(true)
}

func (i *Impl) InitSchema() {
	i.DB.AutoMigrate(&Idol{})
}
