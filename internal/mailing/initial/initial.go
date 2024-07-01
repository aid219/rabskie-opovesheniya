package initial

import (
	"errors"
	"log/slog"
	"rabiKrabi/config"
	"rabiKrabi/internal/mailing"
	"rabiKrabi/internal/mailing/mail"
	"rabiKrabi/internal/mailing/telega"
)

func Init(log *slog.Logger) (map[string]mailing.Messager, error) {
	msgrs := make(map[string]mailing.Messager)
	mes := config.ConfigData.Messangers
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
	if len(msgrs) == 0 {
		log.Error("No messangers found!")
		err := errors.New("no messangers found")
		return nil, err
	}
	log.Info("Collected messangers!")
	return msgrs, nil
}

func InitAllSenders(log *slog.Logger) (map[string]mailing.Messager, error) {
	messangers, err := Init(log)
	if err != nil {
		return nil, err
	}
	for _, i := range messangers {
		i.Init(log)
	}
	log.Info("All messangers init!")
	return messangers, nil
}
