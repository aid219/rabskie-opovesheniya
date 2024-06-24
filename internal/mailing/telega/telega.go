package telega

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

type Telega struct {
	bot       *tgbotapi.BotAPI
	Recipient string
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

func (t *Telega) SetRecepient(rec string) error {
	t.Recipient = rec
	return nil
}

func (t *Telega) Send(mailContent string) error {
	// sdfs := "324234234"
	// chatID, err := strconv.Atoi(sdfs)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// msg := tgbotapi.NewMessage(int64(chatID), mailContent)
	// _, err = t.bot.Send(msg)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// log.Println("Все отправлено!")
	return nil
}
