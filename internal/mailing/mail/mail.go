package mail

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gopkg.in/mail.v2"
)

type Mail struct {
	message    mail.Message
	dialer     mail.Dialer
	Recipients []string
}

func (m *Mail) Init() error {
	mailHeader := "Qummy notifications"
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	m.message = *(mail.NewMessage())
	m.message.SetHeader("From", os.Getenv("MAIL_SENDER"))
	m.message.SetHeader("Subject", mailHeader)
	port, _ := strconv.ParseInt(os.Getenv("MAIL_PORT"), 10, 64)
	s := os.Getenv("MAIL_HOST")
	_ = s
	m.dialer = *(mail.NewDialer(os.Getenv("MAIL_HOST"), int(port), os.Getenv("MAIL_USERNAME"), os.Getenv("MAIL_PASSWORD")))
	m.dialer.StartTLSPolicy = mail.MandatoryStartTLS
	return nil
}

func (t *Mail) SetRecepient(rec string) error {
	t.Recipients = append(t.Recipients, rec)
	return nil
}

func (m *Mail) Send(mailContent string) error {
	for _, i := range m.Recipients {
		m.message.SetHeader("To", i)
		m.message.SetBody("text/plain", mailContent)
		err := m.dialer.DialAndSend(&m.message)
		if err != nil {
			fmt.Println(err)
			// return nil
		}
		fmt.Println(mailContent)
	}
	return nil
}
