package models

import (
	"time"
)

type Chat struct {
	ID string `gorm:"primaryKey"`

	PhraseWordsCount int           // Количество слов в каждой фразе
	OptionsCount     int           // Количество вариантов
	SecondsToVerify  time.Duration // Время, которое дается новому участнику для прохождения проверки
	WelcomeText      string        // Текст-приветствие новым участникам группы в формате HTML

	CreatedAt time.Time
	UpdatedAt time.Time
}

const (
	DefaultSecondsToVerify = time.Second * 60
	DefaultWelcomeTextRU   = "Добро пожаловать! Подтверди, что ты не робот. " +
		"Для этого нужно выбрать фразу, которая изображена на картинке"
	DefaultWelcomeTextEN = "Welcome! Confirm that you are not a robot. " +
		"To do this, select the phrase shown in the picture"
)

func NewChat(chatId string, phraseWordCount int, optionsCount int, langCode string) *Chat {
	var welcomeMsg string
	switch langCode {
	case "ru":
		welcomeMsg = DefaultWelcomeTextRU
	default:
		welcomeMsg = DefaultWelcomeTextEN
	}
	return &Chat{
		ID:               chatId,
		PhraseWordsCount: phraseWordCount,
		OptionsCount:     optionsCount,
		WelcomeText:      welcomeMsg,
		SecondsToVerify:  DefaultSecondsToVerify,
	}
}
