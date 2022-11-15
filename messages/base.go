package messages

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

const (
	NotEnoughChatRights   = "NOT_ENOUGH_CHAT_RIGHTS"
	CantSetWelcomeText    = "CANT_SET_WELCOME_TEXT"
	CantSetVerifySeconds  = "CANT_SET_VERIFY_SECONDS"
	CantSetWordsPerPhrase = "CANT_SET_WORDS_PER_PHRASE"
	CantSetOptionsCount   = "CANT_SET_OPTIONS_COUNT"
	CaptchaSuccess        = "CAPTCHA_SUCCESS"
)

func InitTranslations() {
	initEN()
	initRU()
}

func NewPrinter(langCode string) *message.Printer {
	switch langCode {
	case "ru":
		return message.NewPrinter(language.Russian)
	default:
		return message.NewPrinter(language.English)
	}
}
