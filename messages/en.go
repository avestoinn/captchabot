package messages

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func initEN() {
	_ = message.SetString(language.English, CaptchaSuccess, "You've passed the captcha! Good job!")
	_ = message.SetString(language.English, NotEnoughChatRights, "Only admins can execute bot commands")
	_ = message.SetString(language.English, CantSetWelcomeText, "Cannot set chat's welcome text. Check chat's rights and permission")
	_ = message.SetString(language.English, CantSetVerifySeconds, "Cannot set verify seconds param. Please, check your input is correct. Value has to be more than 10 and less than 360")
	_ = message.SetString(language.English, CantSetWordsPerPhrase, "Cannot set words per phrase param value. Please, check your input is correct. Value has to be more than or equal to 1 and less than or equal to 5")
	_ = message.SetString(language.English, CantSetOptionsCount, "Cannot set options count param value. Please, check your input is correct. Values has to be more than or equal to 3 and less than or equal to 7")
}
