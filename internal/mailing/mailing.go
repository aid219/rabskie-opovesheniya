package mailing

import (
	"encoding/json"
	"log/slog"
	"os"

	"github.com/streadway/amqp"
)

type Message struct {
	Topic string `json:"topic"`
	Body  string `json:"body"`
	HTML  string `json:"html"`
}

type To struct {
	Type      string `json:"type"`
	Recipient string `json:"recipient"`
}

type InData struct {
	To      []To    `json:"to"`
	Message Message `json:"message"`
}

func Receive(log *slog.Logger, ch *amqp.Channel, q *amqp.Queue, senders map[string]Messager) error {
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		log.Error("Error consuming queue : ", err)
		return err
	}
	err = ch.Qos(
		1000,  // prefetchCount: количество сообщений, которые может получить потребитель за один раз
		0,     // prefetchSize
		false, // global
	)
	if err != nil {
		log.Error("Error setting qos : ", err)
		return err
	}
	// Бесконечно ждем сообщения в горутине
	// Горутина для получения сообщений

	go func() {

		defer ch.Close()
		for d := range msgs {
			parsedData, err := Parsing(log, d.Body)
			_ = parsedData
			if err != nil {
				log.Error("Error parsing JSON from rabbit: ", err)
				file, _ := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				file.WriteString(string(d.Body) + "\n")
				file.Close()
				d.Reject(false)
				continue
			}

			err = Send(log, parsedData, senders)
			if err != nil {
				log.Error("Error send message: ", err)
				d.Reject(true)
				continue
			} else {
				d.Ack(true)
			}

			// Отправляем содержимое сообщения через канал out
		}
	}()

	return nil
}

func Parsing(log *slog.Logger, income []byte) (InData, error) {
	var inD InData
	err := json.Unmarshal(income, &inD)
	if err != nil {
		log.Error("Error parsing JSON: ", err)
		return inD, err
	}
	return inD, nil
}
func Send(log *slog.Logger, income InData, senders map[string]Messager) error {
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
