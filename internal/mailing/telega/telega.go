package telega

import (
	"log/slog"
	"rabiKrabi/config"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (t *Telega) Init(log *slog.Logger) error {
	// Получаем токен Telegram бота из конфигурации
	botToken := config.ConfigData.Telegram.Token
	// Создаем новый экземпляр Telegram бота
	bot, err := tgbotapi.NewBotAPI(botToken)
	// Сохраняем экземпляр бота в поле структуры
	t.bot = bot
	if err != nil {
		// Если произошла ошибка при создании бота, логируем её и возвращаем ошибку
		log.Error("Error initializing Telegram bot", err)
		return err
	}
	// Если инициализация прошла успешно, логируем успех и возвращаем nil
	log.Info("Telegram bot initialized successfully!")
	return nil
}

func (t *Telega) Send(log *slog.Logger, chatID string, mailTopic string, mailContent string, timeOut uint16) error {
	// Преобразуем chatID в int64
	chatIDInt, err := strconv.ParseInt(chatID, 10, 64)
	if err != nil {
		// Если возникла ошибка при преобразовании chatID, логируем её и возвращаем ошибку
		log.Error("Error converting chatID", err)
		return err
	}

	// Формируем полное сообщение
	fullMessage := mailTopic + "\n\n" + mailContent
	// Создаем новое сообщение для отправки через Telegram бота
	msg := tgbotapi.NewMessage(chatIDInt, fullMessage)

	// Задержка перед отправкой сообщения
	time.Sleep(time.Millisecond * time.Duration(timeOut))

	// Отправляем сообщение через Telegram бота
	_, err = t.bot.Send(msg)
	if err != nil {
		// Если произошла ошибка при отправке сообщения, логируем её и возвращаем ошибку
		log.Error("Error sending message", err)
		return err
	}

	// Если сообщение успешно отправлено, логируем успех
	log.Info("Message sent successfully to chatID: " + chatID)
	return nil
}
