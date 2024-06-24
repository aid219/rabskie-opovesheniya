package initial

import (
	"rabiKrabi/config"
	"rabiKrabi/internal/mailing"
	"rabiKrabi/internal/mailing/mail"
	"rabiKrabi/internal/mailing/telega"
)

func Init() map[string]mailing.Messager {
	msgrs := make(map[string]mailing.Messager)
	mes := config.LoadConfig().Messangers
	for _, i := range mes {
		switch i {
		case "telegram":
			var m mailing.Messager = &telega.Telega{}
			msgrs[i] = m
		case "email":
			var m mailing.Messager = &mail.Mail{}
			msgrs[i] = m
		}
	}
	return msgrs
}

func InitAllSenders() map[string]mailing.Messager {
	messangers := Init()
	for _, i := range messangers {
		i.Init()
	}
	return messangers
}
