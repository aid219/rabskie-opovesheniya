package initial

import (
	"errors"
	"log/slog"
	"rabiKrabi/config"
	"rabiKrabi/internal/mailing"
	"rabiKrabi/internal/mailing/mail"
	"rabiKrabi/internal/mailing/telega"
)

// Init инициализирует мессенджеры на основе конфигурационных данных и возвращает карту мессенджеров и ошибку, если она есть.
func Init(log *slog.Logger) (map[string]mailing.Messager, error) {
	msgrs := make(map[string]mailing.Messager)
	mes := config.ConfigData.Messangers

	// Перебираем конфигурационные данные для мессенджеров
	for _, i := range mes {
		switch i {
		case "telegram":
			var m mailing.Messager = &telega.Telega{} // Создаем экземпляр Telegram бота
			msgrs[i] = m
		case "email":
			var m mailing.Messager = &mail.Mail{} // Создаем экземпляр объекта для отправки почты
			msgrs[i] = m
		}
	}

	// Если список мессенджеров пуст, возвращаем ошибку
	if len(msgrs) == 0 {
		log.Error("No messangers found!")
		err := errors.New("no messangers found")
		return nil, err
	}

	log.Info("Collected messangers!")
	return msgrs, nil
}

// InitAllSenders инициализирует все мессенджеры и возвращает карту мессенджеров и ошибку, если она есть.
func InitAllSenders(log *slog.Logger) (map[string]mailing.Messager, error) {
	// Инициализируем мессенджеры с помощью функции Init
	messangers, err := Init(log)
	if err != nil {
		return nil, err
	}

	// Вызываем метод Init() для каждого мессенджера
	for _, i := range messangers {
		i.Init(log)
	}

	log.Info("All messangers init!")
	return messangers, nil
}
