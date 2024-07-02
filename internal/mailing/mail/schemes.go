package mail

import "gopkg.in/mail.v2"

// Mail представляет структуру для отправки электронных писем.
type Mail struct {
	message mail.Message // Поле для хранения сообщения письма
	dialer  mail.Dialer  // Поле для хранения Dialer'а для отправки писем
}
