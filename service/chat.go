package service

import (
	"errors"
	"github.com/avestoinn/captchabot/models"
	"gorm.io/gorm"
	"log"
)

func GetChatById(chatId string) *models.Chat {
	var chat = models.Chat{ID: chatId}
	if tx := models.DB.First(&chat); tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			log.Println("Cannot find Chat with such ID. Error: ", tx.Error.Error())
		} else {
			log.Println("Error occurred while trying to find Chat. Error: ", tx.Error.Error())
		}
		return nil
	}
	return &chat
}
