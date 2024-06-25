package initial

import (
	"log/slog"
	"rabiKrabi/globals"
	"rabiKrabi/internal/mailing"
	"rabiKrabi/internal/mailing/mail"
	"rabiKrabi/internal/mailing/telega"
)

func Init(log *slog.Logger) map[string]mailing.Messager {
	msgrs := make(map[string]mailing.Messager)
	mes := globals.ConfigData.Messangers
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
	log.Info("Collected messangers: %v", msgrs)
	return msgrs
}

func InitAllSenders(log *slog.Logger) map[string]mailing.Messager {
	messangers := Init(log)
	for _, i := range messangers {
		i.Init(log)
	}
	log.Info("All messangers init: %v", messangers)
	return messangers
}
