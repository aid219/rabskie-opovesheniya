package initial

import (
	"rabiKrabi/config"
	"rabiKrabi/internal/mailing"
	"rabiKrabi/internal/mailing/mail"
	"rabiKrabi/internal/mailing/telega"
)

func Init() []mailing.Messager {
	var msgrs []mailing.Messager
	mes := config.LoadConfig().Messangers
	for _, i := range mes {
		switch i {
		case "telegram":
			var m mailing.Messager = &(telega.Telega{})
			msgrs = append(msgrs, m)
		case "email":
			var m mailing.Messager = &(mail.Mail{})
			msgrs = append(msgrs, m)
		}
	}
	return msgrs
}

func setRec(mes []mailing.Messager, sender string, data string) {
	switch sender {
	case "telegram":
		mes[0].SetRecepient(data)
	case "email":
		mes[0].SetRecepient(data)
	}
}

func InitAllSenders() []mailing.Messager {
	messangers := Init()
	for _, i := range messangers {
		i.Init()
	}
	return messangers
}

func SendAll(rec []mailing.Messager, data string) error {
	for _, i := range rec {
		i.Send(data)
	}
	return nil
}
