package rabbit

import (
	"log/slog"

	"github.com/streadway/amqp"
)

// Init устанавливает соединение с RabbitMQ, создает канал и объявляет очередь.
// Возвращает канал (ch) для работы с сообщениями, объект очереди (q) и ошибку (если есть).
// Входные параметры:
// - log: Логгер для записи информации об ошибках и действиях.
// - host: Адрес хоста RabbitMQ, к которому нужно подключиться.
// - queueName: Имя очереди, которую нужно объявить.
// Возвращает:
// - *amqp.Channel: Канал для отправки и получения сообщений.
// - *amqp.Queue: Указатель на объект очереди, объявленной в RabbitMQ.
// - error: Ошибка, возникающая при неудачном подключении, создании канала или объявлении очереди.
func Init(log *slog.Logger, host string, queueName string) (*amqp.Channel, *amqp.Queue, error) {
	// Устанавливаем соединение с RabbitMQ
	conn, err := amqp.Dial(host)
	if err != nil {
		log.Error("Error dialing broker : ", err)
		return nil, nil, err
	}

	// Создаем канал для работы с RabbitMQ
	ch, err := conn.Channel()
	if err != nil {
		log.Error("Error creating channel : ", err)
		return nil, nil, err
	}

	// Объявляем очередь в RabbitMQ
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
