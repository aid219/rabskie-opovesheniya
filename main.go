package main

import (
	"log"

	"gopkg.in/mail.v2"
)

// ///sdasddasd
func main() {
	// Создаем новое сообщение
	m := mail.NewMessage()

	// Устанавливаем отправителя
	m.SetHeader("From", "qum.my@yandex.ru")
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter text: ")
	// text, _ := reader.ReadString('\n')
	// fmt.Println(text)
	// Устанавливаем получателя
	m.SetHeader("To", "ahn.ngx@gmail.com")

	// Устанавливаем тему письма
	m.SetHeader("Subject", "Письмо плюшке")

	// Устанавливаем тело письма
	m.SetBody("text/plain", "love love love")

	// Настройка SMTP сервера
	d := mail.NewDialer("smtp.yandex.ru", 587, "qum.my@yandex.ru", "ltpiedglwsdvbomb")
	// m.SetHeader("MIME-Version", "1.0")
	// m.SetHeader("Content-Type", "text/plain; charset=\"UTF-8\"")
	// sender(d, m)
	// Отправка сообщения
	if err := d.DialAndSend(m); err != nil {
		log.Fatal(err)
	}
}

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
