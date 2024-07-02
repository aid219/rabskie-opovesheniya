package mailing

import (
	"encoding/json"
	"log/slog"
	"os"

	"github.com/streadway/amqp"
)

// Receive запускает процесс получения сообщений из очереди RabbitMQ и отправки их через мессенджеры.
// Принимает log для логирования, ch для канала RabbitMQ, q для очереди RabbitMQ и senders для карты мессенджеров.
// Возвращает ошибку, если произошла ошибка при получении или обработке сообщений.
func Receive(log *slog.Logger, ch *amqp.Channel, q *amqp.Queue, senders map[string]Messager) error {
	// Получение сообщений из очереди
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

	// Установка настроек QoS для канала
	err = ch.Qos(
		1000,  // prefetchCount: количество сообщений, которые может получить потребитель за один раз
		0,     // prefetchSize
		false, // global
	)
	if err != nil {
		log.Error("Error setting qos : ", err)
		return err
	}

	// Горутина для получения и обработки сообщений
	go func() {
		defer ch.Close()

		for d := range msgs {
			// Парсинг полученных данных
			parsedData, err := Parsing(log, d.Body)
			if err != nil {
				log.Error("Error parsing JSON from rabbit: ", err)
				// Запись некорректных данных в файл и отклонение сообщения
				file, _ := os.OpenFile("./logs/IncorrectRabbitData.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				file.WriteString(string(d.Body) + "\n")
				file.Close()
				d.Reject(false)
				continue
			}

			// Отправка данных через мессенджеры
			err = Send(log, parsedData, senders)
			if err != nil {
				log.Error("Error send message: ", err)
				d.Reject(true)
				continue
			} else {
				d.Ack(true)
			}
		}
	}()

	return nil
}

// Parsing выполняет парсинг JSON из байтового массива в структуру InData.
// Принимает log для логирования и income - входные данные для парсинга.
// Возвращает распарсенные данные типа InData и ошибку, если произошла ошибка парсинга.
func Parsing(log *slog.Logger, income []byte) (InData, error) {
	var inD InData
	err := json.Unmarshal(income, &inD)
	if err != nil {
		log.Error("Error parsing JSON: ", err)
	}
	return inD, err
}

// Send отправляет сообщения через мессенджеры, определенные в senders, на основе данных income.
// Принимает log для логирования, income с данными для отправки и senders - мапу мессенджеров.
// Возвращает ошибку, если произошла ошибка при отправке сообщения.
func Send(log *slog.Logger, income InData, senders map[string]Messager) error {
	if len(income.To) > 0 {
		for _, i := range income.To {
			switch i.Type {
			case "email":
				if income.Message.HTML == "" {
					log.Info("Send email" + "/" + i.Recipient + "/" + income.Message.Topic + "/" + income.Message.Body)
					senders[i.Type].Send(log, i.Recipient, income.Message.Topic, income.Message.Body, 500)
				} else {
					log.Info("Send email" + "/" + i.Recipient + "/" + income.Message.Topic + "/" + income.Message.HTML)
					senders[i.Type].Send(log, i.Recipient, income.Message.Topic, income.Message.HTML, 500)
				}

			case "telegram":
				log.Info("Send telegram" + "/" + i.Recipient + "/" + income.Message.Topic + "/" + income.Message.Body)
				senders[i.Type].Send(log, i.Recipient, income.Message.Topic, income.Message.Body, 500)
			}
		}
	} else {
		log.Info("Nothing to send")
		return nil
	}
	log.Info("All mail sent successfully!")
	return nil
}
