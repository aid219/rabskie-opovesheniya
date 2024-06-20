package mail

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gopkg.in/mail.v2"
)

type MailAgent struct {
	message mail.Message
	dialer  mail.Dialer
}

func InitMailAgent() (error, *MailAgent) {
	var agent MailAgent

	mailHeader := "Qummy notifications"

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	m := mail.NewMessage()
	m.SetHeader("From", os.Getenv("MAIL_SENDER"))
	m.SetHeader("Subject", mailHeader)

	port, _ := strconv.ParseInt(os.Getenv("MAIL_PORT"), 10, 64)
	s := os.Getenv("MAIL_HOST")
	_ = s
	d := mail.NewDialer(os.Getenv("MAIL_HOST"), int(port), os.Getenv("MAIL_USERNAME"), os.Getenv("MAIL_PASSWORD"))
	d.StartTLSPolicy = mail.MandatoryStartTLS
	agent = MailAgent{message: *m, dialer: *d}

	return nil, &agent
}

func SendMail(agent *MailAgent, mailRecipient string, mailContent string) error {

	agent.message.SetHeader("To", mailRecipient)
	agent.message.SetBody("text/plain", mailContent)

	err := agent.dialer.DialAndSend(&agent.message)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
