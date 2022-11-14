package models

import (
	"github.com/avestoinn/captchabot/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func SetupDb() {
	var err error
	DB, err = gorm.Open(sqlite.Open(config.Config.Database.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Cannot establish a database connection. Error: %v", err.Error())
	}
	err = DB.AutoMigrate(&Chat{})
}
