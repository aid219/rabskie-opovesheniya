package main

import (
	"log"
	"rabiKrabi/internal/mailing/initial"
	"rabiKrabi/internal/rabbit"
)

// 1150762777 tg
func main() {
	a, b, err := rabbit.Init("amqp://guest:guest@localhost:5672/", "rabbitsLoveAnal")
	if err != nil {
		log.Fatal(err)
	}
	send := initial.InitAllSenders()
	aa, _ := rabbit.Receive(a, b)
	for {
		in := <-aa
		if len(in.To) > 0 {
			go rabbit.Send(in, send)
		}

	}
	// senders := initial.InitAllSenders()
	// senders[0].SetRecepient("659623386")
	// go sss(senders)
	// senders[0].SetRecepient("1150762777")
	// senders[0].SetRecepient("aid219@mail.ru")
	// senders[0].SetRecepient("aid219@yandex.ru")

	// initial.SendAll(senders, "vse gorit")

}

// err = mail.SendMail(agent, "aid219@mail.ru", "pisochka")

// if err != nil {
// 	log.Fatal(err)
// }
// err, agent := telega.InitTelegramAgent()
// if err != nil {
// 	log.Fatal(err)
// }
// var id int64 = 1150762777
// err = telega.SendMail(agent, id, "nu ti i padla")
// if err != nil {
// 	log.Fatal(err)
// }

// import (
// 	"fmt"
// 	"os"

// 	"github.com/joho/godotenv"
// 	twilio "github.com/twilio/twilio-go"
// 	openapi "github.com/twilio/twilio-go/rest/api/v2010"
// )

// func main() {
// 	if err := godotenv.Load(); err != nil {
// 		fmt.Println("Error loading .env file:", err)
// 		return
// 	}

// 	client := twilio.NewRestClient()
// 	// twilio.Client()
// 	params := &openapi.CreateMessageParams{}
// 	params.SetTo(os.Getenv("TO_PHONE_NUMBER"))
// 	params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER"))
// 	params.SetPathAccountSid(os.Getenv("TWILIO_ACCOUNT_SID"))
// 	params.SetBody("Hello from Golang!")

// 	_, err := client.Api.CreateMessage(params)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	} else {
// 		fmt.Println("SMS sent successfully!")
// 	}
// }

// import (
// 	"log"
// 	"os"
// 	"strconv"

// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
// 	"github.com/joho/godotenv"
// )

// func main() {
// 	// Загрузить переменные окружения из файла .env
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
// 	chatID, err := strconv.ParseInt(os.Getenv("TELEGRAM_CHAT_ID"), 10, 64)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// Создать новый бот
// 	bot, err := tgbotapi.NewBotAPI(botToken)
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	// Создать новое сообщение
// 	msg := tgbotapi.NewMessage(chatID, "hello suchara")

// 	// Отправить сообщение
// 	_, err = bot.Send(msg)
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	log.Println("Сообщение успешно отправлено!")
// }

// import (
// 	"log"

// 	"gopkg.in/mail.v2"
// )

// func main() {
// 	token := "7191966794:AAEOt99eIGICjTcB0Fobz8hWMCD_2uU-gfw"
// 	// Создаем новое сообщение
// 	m := mail.NewMessage()

// 	// Устанавливаем отправителя
// 	m.SetHeader("From", "qum.my@yandex.ru")
// 	// reader := bufio.NewReader(os.Stdin)
// 	// fmt.Print("Enter text: ")
// 	// text, _ := reader.ReadString('\n')
// 	// fmt.Println(text)
// 	// Устанавливаем получателя
// 	m.SetHeader("To", "ahn.ngx@gmail.com")

// 	// Устанавливаем тему письма
// 	m.SetHeader("Subject", "Письмо плюшке")

// 	// Устанавливаем тело письма
// 	m.SetBody("text/plain", "love love love")

// 	// Настройка SMTP сервера
// 	d := mail.NewDialer("smtp.yandex.ru", 587, "qum.my@yandex.ru", "ltpiedglwsdvbomb")
// 	// m.SetHeader("MIME-Version", "1.0")
// 	// m.SetHeader("Content-Type", "text/plain; charset=\"UTF-8\"")
// 	// sender(d, m)
// 	// Отправка сообщения
// 	if err := d.DialAndSend(m); err != nil {
// 		log.Fatal(err)
// 	}
// }

// func sender(dd *mail.Dialer, mm *mail.Message) {
// 	i := " love"

// 	for {
// 		i += i
// 		mm.SetBody("text/plain", "love love love"+i)
// 		if err := dd.DialAndSend(mm); err != nil {
// 			log.Fatal(err)
// 		}
// 		time.Sleep(5 * time.Second)
// 	}

// }
