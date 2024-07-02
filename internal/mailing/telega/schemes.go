package telega

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// Telega представляет структуру для работы с Telegram ботом.
type Telega struct {
	bot *tgbotapi.BotAPI // Поле для хранения экземпляра Telegram бота
}
