package bot

import (
	"fmt"
	"github.com/avestoinn/captchabot/messages"
	"github.com/avestoinn/captchabot/models"
	"github.com/avestoinn/captchabot/service"
	tele "gopkg.in/telebot.v3"
	"strconv"
)

func GroupMessageOnlyMiddleware(f tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		var txtMessage string
		switch c.Sender().LanguageCode {
		case "ru":
			txtMessage = "Привет, друг! Я умею бороться с роботами в группах и отличать их от людей. " +
				"Добавь меня в группу, чтобы я заработал!"
		default:
			txtMessage = "Hello, friend! I know how to fight robots in groups and distinguish them from humans. " +
				"Add me to the group to make it work!"
		}
		if c.Chat().Type == tele.ChatPrivate {
			return c.Send(txtMessage)
		} else if c.Chat().Type == tele.ChatSuperGroup || c.Chat().Type == tele.ChatGroup {
			chat := service.GetChatById(fmt.Sprintf("%v", c.Chat().ID))
			if chat == nil {
				chat = models.NewChat(strconv.Itoa(int(c.Chat().ID)), 2, 3, c.Sender().LanguageCode)
				models.DB.Save(&chat)
			}
			c.Set("contextChat", chat)
			return f(c)
		}
		return nil
	}
}

func ChatAdminOnlyMiddleware(f tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		sender := c.Sender()
		p := messages.NewPrinter(sender.LanguageCode)

		// Getting chat admins
		adminMembers, err := c.Bot().AdminsOf(c.Chat())
		if err != nil {
			return err
		}

		// Returns handler func only if sender is a chat admin
		for _, admin := range adminMembers {
			if sender.ID == admin.User.ID {
				return f(c)
			}
		}
		return c.Send(p.Sprintf(messages.NotEnoughChatRights))
	}
}
