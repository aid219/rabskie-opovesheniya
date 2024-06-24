package rabbit

// {
// 	"to": [
// 	  {"type": "email", "recipient": "aid219@mail.ru"},
// 	  {"type": "email", "recipient": "Aid219@yandex.ru"}
// 	],
// 	"message": {
// 	  "topic":"wuwu",
// 	  "body": "hi bomzh",
// 	  "HTML": ""
// 	}
//   }

import (
	"encoding/json"
	"fmt"
	"log"
	"rabiKrabi/internal/mailing"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Init(host string, queueName string) (*amqp.Channel, *amqp.Queue, error) {
	conn, err := amqp.Dial(host)
	failOnError(err, "Failed to connect to RabbitMQ")
	// defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	// defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")
	return ch, &q, nil
}

func Receive(ch *amqp.Channel, q *amqp.Queue) (chan InData, error) {
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	// Бесконечно ждем сообщения в горутине
	out := make(chan InData)

	// Горутина для получения сообщений
	go func() {
		defer close(out)
		defer ch.Close()
		for d := range msgs {
			// log.Printf("Received a message: %s", d.Body)
			parsedData, err := RabPars(d.Body)
			if err != nil {
				log.Fatalf("Error parsing JSON: %v", err)
			}
			out <- parsedData // Отправляем содержимое сообщения через канал out
		}
	}()
	return out, nil
}

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

func RabPars(income []byte) (InData, error) {
	var inD InData
	err := json.Unmarshal(income, &inD)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}
	return inD, nil
}

func Send(income InData, senders map[string]mailing.Messager) error {

	mailling := income
	if len(mailling.To) > 0 {
		for _, i := range mailling.To {
			switch i.Type {

			case "email":
				if mailling.Message.HTML == "" {
					fmt.Println(i.Recipient, mailling.Message.Topic, mailling.Message.Body)
					senders[i.Type].Send(i.Recipient, mailling.Message.Topic, mailling.Message.Body)
				} else {
					fmt.Println(i.Recipient, mailling.Message.Topic, mailling.Message.HTML)
					senders[i.Type].Send(i.Recipient, mailling.Message.Topic, mailling.Message.HTML)
				}

			case "telegram":

				senders[i.Type].Send(i.Recipient, mailling.Message.Topic, mailling.Message.Body)

			}

		}
	}
	return nil
}
