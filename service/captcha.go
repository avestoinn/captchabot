package service

import (
	"bytes"
	"github.com/avestoinn/text2img"
	"github.com/disintegration/imaging"
	"github.com/tjarratt/babble"
	"golang.org/x/image/colornames"
	"image/jpeg"
	"log"
	"math/rand"
	"time"
)

type Captcha struct {
	Phrases []Phrase
	Image   *bytes.Buffer

	phraseWordCount int
}

type Phrase struct {
	Text      string
	IsCorrect bool
}

func NewCaptcha(optionsCount int, phraseWordCount int) *Captcha {
	if optionsCount <= 0 {
		optionsCount = 3
	}

	captcha := &Captcha{
		Phrases:         []Phrase{},
		Image:           nil,
		phraseWordCount: phraseWordCount,
	}

	for i := 0; i < optionsCount; i++ {
		ph := Phrase{
			Text:      generatePhrase(captcha.phraseWordCount),
			IsCorrect: false,
		}
		captcha.Phrases = append(captcha.Phrases, ph)
	}

	rand.Seed(time.Now().Unix())
	randPhIndex := rand.Intn(len(captcha.Phrases))
	captcha.Phrases[randPhIndex].IsCorrect = true
	captcha.Image = generateImage(captcha.Phrases[randPhIndex].Text)

	return captcha
}

func generateImage(text string) *bytes.Buffer {
	d, err := text2img.NewDrawer(text2img.Params{
		BackgroundColor: colornames.Darkslategrey,
		TextColor:       colornames.White,
		FontPath:        "fonts/ubuntu.ttf",
	})
	if err != nil {
		log.Println("Cannot initialze NewDrawer. Error: ", err.Error())
		return nil
	}

	img, err := d.Draw(text)
	if err != nil {
		log.Println("Cannot draw text. Error: ", err.Error())
		return nil
	}
	blurred := imaging.Blur(img, 6)
	buf := bytes.NewBuffer(*new([]byte))
	err = jpeg.Encode(buf, blurred, &jpeg.Options{Quality: 100})
	if err != nil {
		log.Println("Cannot encode jpeg image. Error: ", err.Error())
		return nil
	}

	return buf
}

func generatePhrase(wordsCount int) string {
	b := babble.NewBabbler()
	b.Separator = " "

	if wordsCount <= 0 {
		b.Count = 1
	} else {
		b.Count = wordsCount
	}

	return b.Babble()
}
