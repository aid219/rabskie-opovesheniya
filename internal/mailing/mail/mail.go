package mail

import (
	"log/slog"
	"rabiKrabi/config"
	"time"

	"gopkg.in/mail.v2"
)

func (m *Mail) Init(log *slog.Logger) error {

	m.message = *(mail.NewMessage())
	m.message.SetHeader("From", config.ConfigData.Mail.Sender)
	m.dialer = *(mail.NewDialer(config.ConfigData.Mail.Host, int(config.ConfigData.Mail.Port), config.ConfigData.Mail.Username, config.ConfigData.Mail.Password))
	m.dialer.StartTLSPolicy = mail.MandatoryStartTLS
	log.Info("mail sender init successfully!")
	return nil
}

func (m *Mail) Send(log *slog.Logger, mailRecipient string, mailTopic string, mailContent string, timeOut uint16) error {
	time.Sleep(time.Millisecond * time.Duration(timeOut))
	m.message.SetHeader("To", mailRecipient)
	m.message.SetHeader("Subject", mailTopic)
	m.message.SetBody("text/plain", mailContent)
	err := m.dialer.DialAndSend(&m.message)

	if err != nil {
		log.Error("Error sending mail", err)
	}
	log.Info("mail sent successfully!: " + mailRecipient)
	return err
}
