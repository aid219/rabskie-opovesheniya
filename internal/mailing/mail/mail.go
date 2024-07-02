package mail

import (
	"log/slog"
	"rabiKrabi/config"
	"time"

	"gopkg.in/mail.v2"
)

// Init инициализирует отправителя электронной почты на основе конфигурационных данных.
func (m *Mail) Init(log *slog.Logger) error {
	// Создаем новое сообщение
	m.message = *(mail.NewMessage())
	// Устанавливаем отправителя письма из конфигурации
	m.message.SetHeader("From", config.ConfigData.Mail.Sender)
	// Создаем новый Dialer для отправки писем, используя данные из конфигурации
	m.dialer = *(mail.NewDialer(config.ConfigData.Mail.Host, int(config.ConfigData.Mail.Port), config.ConfigData.Mail.Username, config.ConfigData.Mail.Password))
	// Устанавливаем политику обязательного использования StartTLS
	m.dialer.StartTLSPolicy = mail.MandatoryStartTLS
	log.Info("Mail sender initialized successfully!")
	return nil
}

// Send отправляет электронное письмо с заданными параметрами.
func (m *Mail) Send(log *slog.Logger, mailRecipient string, mailTopic string, mailContent string, timeOut uint16) error {
	// Задержка перед отправкой письма (если указан таймаут)
	time.Sleep(time.Millisecond * time.Duration(timeOut))
	// Устанавливаем получателя письма
	m.message.SetHeader("To", mailRecipient)
	// Устанавливаем тему письма
	m.message.SetHeader("Subject", mailTopic)
	// Устанавливаем тело письма
	m.message.SetBody("text/plain", mailContent)
	// Отправляем письмо через Dialer
	err := m.dialer.DialAndSend(&m.message)

	// Если произошла ошибка при отправке, логируем её
	if err != nil {
		log.Error("Error sending mail", err)
	}
	log.Info("Mail sent successfully to: " + mailRecipient)
	return err
}
