package bot

import (
	"github.com/avestoinn/captchabot/config"
	tele "gopkg.in/telebot.v3"
	"log"
	"time"
)

var Bot *tele.Bot

func Run() {
	pref := tele.Settings{
		Token:  config.Config.Bot.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	var err error
	Bot, err = tele.NewBot(pref)
	if err != nil {
		log.Fatalf("Cannot run bot. Error: %v", err.Error())
	}

	// Default stable routes
	Bot.Use(GroupMessageOnlyMiddleware)
	Bot.Handle(tele.OnUserJoined, SendCaptcha)
	Bot.Handle(&tele.Btn{Unique: "captchaClicked"}, OnCaptchaClicked)

	// Works as a slug to make middlewares work
	Bot.Handle(tele.OnText, func(c tele.Context) error {
		return nil
	})

	// Pre-release support
	// TODO: Добавить поддержку этих команд. На данный момент бот эти команды не поддерживает
	if err = Bot.SetCommands([]tele.Command{
		{Text: "welcome_text", Description: "<text:string> - Set chat's welcome text / Изменить приветственное сообщение"},
		{Text: "verify_seconds", Description: "<seconds:int> - Set seconds to pass captcha / Изменить время прохождения капчи"},
		{Text: "options_count", Description: "<amount:int> - Set options amount per a captcha / Изменить кол-во вариантов для каждой капчи"},
		{Text: "words_per_phrase", Description: "<amount:int> - Set words amount per a phrase / Изменить кол-во слов в каждой фразе (варианте) капчи"},
	}, &tele.CommandScope{Type: tele.CommandScopeAllChatAdmin}); err != nil {
		log.Println("Cannot set bot commands. Error: ", err)
	}

	Bot.Start()
}
