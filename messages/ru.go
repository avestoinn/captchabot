package messages

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func initRU() {
	_ = message.SetString(language.Russian, CaptchaSuccess, "Ты прошел проверку! Молодец, человек!")
	_ = message.SetString(language.Russian, NotEnoughChatRights, "Только админы чата могут вызывать команды бота")
	_ = message.SetString(language.Russian, CantSetWelcomeText, "Невозможно изменить приветственное сообщение. Проверьте настройки прав группы.")
	_ = message.SetString(language.Russian, CantSetVerifySeconds, "Невозможно изменить значение параметра секунд для прохождения капчи. Проверьте, что введенное значение правильное. Значение не может быть меньше 10 и больше 360")
	_ = message.SetString(language.Russian, CantSetWordsPerPhrase, "Невозможно изменить значение параметра слов на фразу. Проверьте, что введенное значение правильное. Значение не может быть меньше 1 и больше 5")
	_ = message.SetString(language.Russian, CantSetOptionsCount, "Невозможно изменить значение параметра фраз на капчу. Проверьте, что введенное значение правильное. Значение не может быть меньше 3 и больше 7")
}
