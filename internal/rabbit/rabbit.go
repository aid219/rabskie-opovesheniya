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
	"log/slog"

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
