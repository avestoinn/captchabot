package service

import (
	"errors"
	"github.com/avestoinn/captchabot/models"
	"gorm.io/gorm"
	"html"
	"log"
	"time"
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

func SetChatWelcomeText(chatId, welcomeText string) (ok bool, err error) {
	escaped := html.EscapeString(welcomeText)
	if tx := models.DB.Model(&models.Chat{}).Where("id = ?", chatId).Update("welcome_text", escaped); tx.Error != nil {
		return false, err
	} else {
		return true, nil
	}
}

func SetChatSecondsToVerify(chatId string, seconds int) (ok bool, err error) {
	if seconds < 10 || seconds > 360 {
		return false, err
	}
	if tx := models.DB.Model(&models.Chat{}).Where("id = ?", chatId).Update("seconds_to_verify", time.Duration(seconds)*time.Second); tx.Error != nil {
		return false, err
	} else {
		return true, nil
	}
}

func SetChatWordsPerPhrase(chatId string, words int) (ok bool, err error) {
	if words < 1 || words > 5 {
		return false, err
	}
	if tx := models.DB.Model(&models.Chat{}).Where("id = ?", chatId).Update("phrase_words_count", words); tx.Error != nil {
		return false, err
	} else {
		return true, nil
	}
}

func SetChatOptionsCount(chatId string, count int) (ok bool, err error) {
	if count < 3 || count > 7 {
		return false, err
	}
	if tx := models.DB.Model(&models.Chat{}).Where("id = ?", chatId).Update("options_count", count); tx.Error != nil {
		return false, err
	} else {
		return true, nil
	}
}
