package mail

import "gopkg.in/mail.v2"

type Mail struct {
	message mail.Message
	dialer  mail.Dialer
}
