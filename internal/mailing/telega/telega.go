package telega

import (
	"log/slog"
	"rabiKrabi/globals"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (t *Telega) Init(log *slog.Logger) error {
	botToken := globals.ConfigData.Telegram.Token
	bot, err := tgbotapi.NewBotAPI(botToken)
	t.bot = bot
	if err != nil {
		log.Error("Error init bot", err)
		return err
	}
	log.Info("telegram sender init successfully!")
	return nil
}

func (t *Telega) Send(log *slog.Logger, mailRecipient string, mailTopic string, mailContent string) error {

	chatID, err := strconv.Atoi(mailRecipient)
	if err != nil {
		log.Error("Error convert chatID", err)
		return err
	}
	fullMessage := mailTopic + "\n\n" + mailContent
	msg := tgbotapi.NewMessage(int64(chatID), fullMessage)
	_, err = t.bot.Send(msg)
	if err != nil {
		log.Error("Error sending message", err)
		return err
	}
	log.Info("message sent successfully!: " + mailRecipient)
	return nil
}
