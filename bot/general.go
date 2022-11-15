package bot

import (
	"fmt"
	"github.com/avestoinn/captchabot/messages"
	"github.com/avestoinn/captchabot/models"
	"github.com/avestoinn/captchabot/service"
	tele "gopkg.in/telebot.v3"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"
)

func getUserMention(userId int64, name string) string {
	return fmt.Sprintf("<a href='tg://user?id=%v'>%v</a>", userId, name)
}

// Here we store users that have passed verification process
var passedUsers = sync.Map{}

func SendCaptcha(c tele.Context) error {
	ctxChat := c.Get("contextChat")
	chat := ctxChat.(*models.Chat)

	// Restricting user until he pass the verification process
	newChatMember := tele.ChatMember{User: c.Message().UserJoined, RestrictedUntil: tele.Forever(),
		Rights: tele.Rights{CanSendMessages: false}}
	err := c.Bot().Restrict(c.Chat(), &newChatMember)
	if err != nil {
		log.Println("Cannot restrict chat member. Error: ", err.Error())
	}

	// Getting a new captcha
	captcha := service.NewCaptcha(chat.OptionsCount, chat.PhraseWordsCount)
	i := &tele.Photo{File: tele.FromReader(captcha.Image), Caption: chat.WelcomeText + "\n\n" +
		getUserMention(c.Sender().ID, c.Sender().FirstName)}

	// Generating a reply markup
	var rows []tele.Row
	markup := &tele.ReplyMarkup{}
	for _, ph := range captcha.Phrases {
		var d = markup.Data(ph.Text, "captchaClicked", fmt.Sprintf("%v|", c.Sender().ID))
		if ph.IsCorrect {
			d.Data += fmt.Sprintf("%v", "correct")
		} else {
			d.Data += strings.SplitAfter(ph.Text, " ")[0]
		}
		rows = append(rows, markup.Row(d))
	}
	markup.Inline(rows...)

	// Sending a prepared message
	msg, _ := c.Bot().Send(c.Chat(), i, markup, tele.ModeHTML)

	// Waiting for user to pass the captcha, otherwise ban him
	time.AfterFunc(chat.SecondsToVerify, func() {
		lastChatId, passed := passedUsers.Load(c.Sender().ID)
		if !passed || lastChatId.(int64) != c.Chat().ID {
			if err := c.Bot().Delete(msg); err != nil {
				log.Printf("Cannot delete message! MsgId %v , ChatId %v , UserId %v",
					msg.ID, c.Chat().ID, c.Sender().ID)
			}

			chatMember := tele.ChatMember{User: c.Message().UserJoined, RestrictedUntil: tele.Forever()}
			err = c.Bot().Ban(c.Chat(), &chatMember)
			if err != nil {
				log.Println("Cannot ban chat member. Error: ", err.Error())
			}
		}
	})

	return nil
}

func OnCaptchaClicked(c tele.Context) error {
	p := messages.NewPrinter(c.Sender().LanguageCode)

	data := strings.Split(c.Data(), "|")
	dsnUserId, _ := strconv.ParseInt(data[0], 10, 64)
	optionName := data[1]

	if c.Sender().ID != dsnUserId {
		return c.Respond(&tele.CallbackResponse{})
	}

	if optionName != "correct" {
		_ = c.Bot().Delete(c.Message())
		chatMember := tele.ChatMember{User: c.Sender(), RestrictedUntil: tele.Forever()}
		return c.Bot().Ban(c.Chat(), &chatMember)
	}

	passedUsers.Store(c.Sender().ID, c.Chat().ID)
	log.Printf("User: %v passed the challenge in chat: %v", c.Sender().ID, c.Chat().ID)
	newChatMember := tele.ChatMember{User: c.Sender(), RestrictedUntil: tele.Forever(),
		Rights: tele.Rights{CanSendMessages: true}}
	err := c.Bot().Promote(c.Chat(), &newChatMember)
	if err != nil {
		log.Println("Cannot promote chat member. Error: ", err.Error())
	}

	_ = c.Respond(&tele.CallbackResponse{Text: p.Sprintf(messages.CaptchaSuccess),
		ShowAlert: true})

	if err = c.Bot().Delete(c.Message()); err != nil {
		log.Println("Cannot delete message after successful captcha. Error: ", err.Error())
	}
	return nil
}
