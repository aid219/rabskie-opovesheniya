package telega

import (
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

type Telega struct {
	bot *tgbotapi.BotAPI
}

func (t *Telega) Init() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	botToken := os.Getenv("TG_TOKEN")
	bot, err := tgbotapi.NewBotAPI(botToken)
	t.bot = bot
	if err != nil {
		log.Panic(err)
	}
	return nil
}

func (t *Telega) Send(mailRecipient string, mailTopic string, mailContent string) error {

	chatID, err := strconv.Atoi(mailRecipient)
	if err != nil {
		log.Panic(err)
	}
	fullMessage := mailTopic + "\n\n" + mailContent
	msg := tgbotapi.NewMessage(int64(chatID), fullMessage)
	_, err = t.bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Все отправлено!")
	return nil
}
