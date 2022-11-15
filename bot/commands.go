package bot

import (
	"fmt"
	"github.com/avestoinn/captchabot/messages"
	"github.com/avestoinn/captchabot/service"
	tele "gopkg.in/telebot.v3"
	"log"
	"strconv"
)

func SetWelcomeText(c tele.Context) error {
	p := messages.NewPrinter(c.Sender().LanguageCode)
	welcomeText := c.Message().Payload

	ok, err := service.SetChatWelcomeText(fmt.Sprintf("%v", c.Chat().ID), welcomeText)
	if !ok || err != nil {
		log.Printf("Cannot set welcome text. Value: %v", welcomeText)
		if err != nil {
			log.Println("Error: ", err.Error())
		}
		return c.Send(p.Sprintf(messages.CantSetWelcomeText))
	}

	return nil
}

func SetVerifySeconds(c tele.Context) error {
	p := messages.NewPrinter(c.Sender().LanguageCode)

	secondsString := c.Message().Payload
	seconds, err := strconv.Atoi(secondsString)
	if err != nil {
		log.Printf("Cannot parse string from value %v. Error: %v", secondsString, err.Error())
		return c.Send(p.Sprintf(messages.CantSetVerifySeconds))
	}

	ok, err := service.SetChatSecondsToVerify(fmt.Sprintf("%v", c.Chat().ID), seconds)
	if !ok || err != nil {
		log.Printf("Cannot set seconds to verify value. Value: %v", seconds)
		if err != nil {
			log.Println("Error: ", err.Error())
		}
		return c.Send(p.Sprintf(messages.CantSetVerifySeconds))
	}

	return nil
}

func SetOptionsCount(c tele.Context) error {
	p := messages.NewPrinter(c.Sender().LanguageCode)

	countStr := c.Message().Payload
	count, err := strconv.Atoi(countStr)
	if err != nil {
		log.Printf("Cannot parse string from value %v. Error: %v", count, err.Error())
		return c.Send(p.Sprintf(messages.CantSetOptionsCount))
	}

	ok, err := service.SetChatOptionsCount(fmt.Sprintf("%v", c.Chat().ID), count)
	if !ok || err != nil {
		log.Printf("Cannot set chat options count value. Value: %v", count)
		if err != nil {
			log.Println("Error: ", err.Error())
		}
		return c.Send(p.Sprintf(messages.CantSetOptionsCount))
	}

	return nil
}

func SetWordsPerPhrase(c tele.Context) error {
	p := messages.NewPrinter(c.Sender().LanguageCode)

	countStr := c.Message().Payload
	count, err := strconv.Atoi(countStr)
	if err != nil {
		log.Printf("Cannot parse string from value %v. Error: %v", count, err.Error())
		return c.Send(p.Sprintf(messages.CantSetWordsPerPhrase))
	}

	ok, err := service.SetChatWordsPerPhrase(fmt.Sprintf("%v", c.Chat().ID), count)
	if !ok || err != nil {
		log.Printf("Cannot set chat words per phrase value. Value: %v", count)
		if err != nil {
			log.Println("Error: ", err.Error())
		}
		return c.Send(p.Sprintf(messages.CantSetWordsPerPhrase))
	}

	return nil
}
