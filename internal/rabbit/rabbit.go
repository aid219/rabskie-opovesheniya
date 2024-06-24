package rabbit

// {"to":[{"type": "email", "recipient":"aid@219@mail.ru"}, {"type": "email", "recipient":"Aid@219@yandex.ru"}], "message":{"text":"sosinator33", "HTML":""}}
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

func Receive(ch *amqp.Channel, q *amqp.Queue) (chan string, error) {
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
	out := make(chan string)

	// Горутина для получения сообщений
	go func() {
		defer close(out)
		defer ch.Close()
		for d := range msgs {
			// log.Printf("Received a message: %s", d.Body)
			out <- string(d.Body) // Отправляем содержимое сообщения через канал out
		}
	}()

	return out, nil
}

type Message struct {
	Text string `json:"text"`
	HTML string `json:"html"`
}

type To struct {
	Type      string `json:"type"`
	Recipient string `json:"recipient"`
}

type InData struct {
	To      []To    `json:"to"`
	Message Message `json:"message"`
}

func RabPars(ch chan string, senders []mailing.Messager) ([]string, string, error) {
	var inD InData
	for {
		bb := <-ch
		// Парсинг JSON строки в структуру
		err := json.Unmarshal([]byte(bb), &inD)
		if len(inD.To) > 0 {
			for _, i := range inD.To {
				if i.Type == "email" {
					senders[1].SetRecepient(i.Recipient)
					senders[1].Send(inD.Message.Text)
				} else if i.Type == "telegram" {
					senders[0].SetRecepient(i.Recipient)
					senders[0].Send(inD.Message.Text)
				}

			}
		}
		fmt.Println(len(inD.To))
		if err != nil {
			log.Fatalf("Error parsing JSON: %v", err)
		}
		// if inD.To.Type[0] == "mail" || inD.To.Type == "telegram" {

		// }
		fmt.Println(inD.Message.Text)
	}

	return nil, "", nil
}

// func send(ch *amqp.Channel, queueName string, message string) {

// 	body := message

// 	err := ch.Publish(
// 		"",        // exchange
// 		queueName, // routing key
// 		false,     // mandatory
// 		false,     // immediate
// 		amqp.Publishing{
// 			ContentType: "text/plain",
// 			Body:        []byte(body),
// 		})

// 	failOnError(err, "Failed to publish a message")
// 	log.Printf(" [x] Sent %s", body)
// }
