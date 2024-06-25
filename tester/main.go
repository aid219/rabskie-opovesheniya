package main

import (
	"strconv"

	"github.com/streadway/amqp"
)

var count int32 = 0
var count2 int32 = 0
var count3 int32 = 0
var count4 int32 = 0

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {

	}
	ch, err := conn.Channel()

	if err != nil {

	}
	q, err := ch.QueueDeclare(
		"rabbitsLoveAnal", // name
		false,             // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)
	_ = q

	for {

		count++
		message := strconv.Itoa(int(count))
		err = ch.Publish(
			"",                // exchange
			"rabbitsLoveAnal", // routing key (queue name)
			false,             // mandatory
			false,             // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(message),
			})
		_ = err
		// time.Sleep(1 * time.Microsecond)
		// fmt.Printf(" [x] Sent %d\n", count)
	}

}
