package mailing

import (
	"encoding/json"
	"log/slog"
)

func Parsing(log *slog.Logger, income []byte) (*InData, error) {
	var inD InData
	err := json.Unmarshal(income, &inD)
	if err != nil {
		log.Error("Error parsing JSON: ", err)
		return nil, err
	}
	return &inD, nil
}

func Send(log *slog.Logger, income []byte, senders map[string]Messager) error {

	parsedData, err := Parsing(log, income)

	if err != nil {
		log.Error("Error parsing JSON in send: ", err)
		return err
	}
	if len(parsedData.To) > 0 {
		for _, i := range parsedData.To {
			switch i.Type {

			case "email":
				if parsedData.Message.HTML == "" {
					log.Info("Send email" + "/" + i.Recipient + "/" + parsedData.Message.Topic + "/" + parsedData.Message.Body)
					err := senders[i.Type].Send(log, i.Recipient, parsedData.Message.Topic, parsedData.Message.Body)
					if err != nil {
						return err
					}

				} else {
					log.Info("Send email" + "/" + i.Recipient + "/" + parsedData.Message.Topic + "/" + parsedData.Message.HTML)
					err := senders[i.Type].Send(log, i.Recipient, parsedData.Message.Topic, parsedData.Message.HTML)
					if err != nil {
						return err
					}
				}

			case "telegram":
				log.Info("Send telegram" + "/" + i.Recipient + "/" + parsedData.Message.Topic + "/" + parsedData.Message.Body)
				err := senders[i.Type].Send(log, i.Recipient, parsedData.Message.Topic, parsedData.Message.Body)
				if err != nil {
					return err
				}

			}

		}
	} else {
		log.Info("Nothing to send")
		return nil
	}
	log.Info("All mail sended successfully!")
	return nil
}
