package telega

import (
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

type Telega struct {
	bot        tgbotapi.BotAPI
	Recipients []string
}

func (t *Telega) Init() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	botToken := os.Getenv("TG_TOKEN")
	bot, err := tgbotapi.NewBotAPI(botToken)
	t.bot = *bot
	if err != nil {
		log.Panic(err)
	}
	return nil
}

func (t *Telega) SetRecepient(rec string) error {
	t.Recipients = append(t.Recipients, rec)
	return nil
}

func (t *Telega) Send(mailContent string) error {
	f := func(rec string) {
		chatID, err := strconv.Atoi(rec)
		if err != nil {
			log.Panic(err)
		}
		msg := tgbotapi.NewMessage(int64(chatID), mailContent)

		_, err = t.bot.Send(msg)
		if err != nil {
			log.Panic(err)
		}
	}
	for _, i := range t.Recipients {
		go f(i)
	}

	log.Println("Все отправлено!")
	return nil
}
