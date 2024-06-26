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
	"os"
	"rabiKrabi/internal/mailing"

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

func ReceiveAndSend(log *slog.Logger, ch *amqp.Channel, q *amqp.Queue, senders map[string]mailing.Messager) error {
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
		return err
	}

	err = ch.Qos(
		100,   // prefetchCount: количество сообщений, которые может получить потребитель за один раз
		0,     // prefetchSize
		false, // global
	)

	// Бесконечно ждем сообщения в горутине
	out := make(chan *amqp.Delivery)

	// Горутина для получения сообщений

	go func() {
		defer close(out)
		defer ch.Close()
		for d := range msgs {
			err := mailing.Send(log, d.Body, senders)
			if err != nil {
				// d.Reject(false)

				file, _ := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				file.WriteString(string(d.Body) + "\n")
				file.Close()

				// log.Error("Error send message : ", err)
				// log.Error(string(d.Body))
				continue
			}
			// d.Ack(false)
		}
	}()

	return nil
}
