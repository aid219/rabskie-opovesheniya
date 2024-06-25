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
	"log/slog"
	"os"

	"github.com/streadway/amqp"
)

func Init(log *slog.Logger, host string, queueName string) (*amqp.Channel, *amqp.Queue, error) {
	conn, err := amqp.Dial(host)
	if err != nil {
		log.Error("Error dialing broker : ", err)
		return nil, nil, err
	}
	ch, err := conn.Channel()

	if err != nil {
		log.Error("Error creating channel : ", err)
		return nil, nil, err
	}

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Error("Error declaring queue : ", err)
		return nil, nil, err
	}
	return ch, &q, nil
}

func Receive(log *slog.Logger, ch *amqp.Channel, q *amqp.Queue) (chan InData, error) {
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		log.Error("Error consuming queue : ", err)
		return nil, err
	}

	// Бесконечно ждем сообщения в горутине
	out := make(chan InData)

	// Горутина для получения сообщений

	go func() {

		defer close(out)
		defer ch.Close()
		for d := range msgs {
			parsedData, err := Parsing(log, d.Body)
			if err != nil {
				log.Error("Error parsing JSON from rabbit: ", err)

				file, _ := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				file.WriteString(string(d.Body) + "\n")
				file.Close()

				continue
			}
			out <- parsedData // Отправляем содержимое сообщения через канал out
		}
	}()

	return out, nil
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
