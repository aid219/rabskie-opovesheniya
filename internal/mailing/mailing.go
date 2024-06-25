package mailing

import (
	"log/slog"
	"rabiKrabi/internal/rabbit"
)

func Send(log *slog.Logger, income rabbit.InData, senders map[string]Messager) error {
	if len(income.To) > 0 {
		for _, i := range income.To {
			switch i.Type {

			case "email":
				if income.Message.HTML == "" {
					log.Info("Send email" + "/" + i.Recipient + "/" + income.Message.Topic + "/" + income.Message.Body)
					senders[i.Type].Send(log, i.Recipient, income.Message.Topic, income.Message.Body)
				} else {
					log.Info("Send email" + "/" + i.Recipient + "/" + income.Message.Topic + "/" + income.Message.HTML)
					senders[i.Type].Send(log, i.Recipient, income.Message.Topic, income.Message.HTML)
				}

			case "telegram":
				log.Info("Send telegram" + "/" + i.Recipient + "/" + income.Message.Topic + "/" + income.Message.Body)
				senders[i.Type].Send(log, i.Recipient, income.Message.Topic, income.Message.Body)

			}

		}
	}
	return nil
}
