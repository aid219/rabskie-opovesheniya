package mailing

type Messager interface {
	Init() error
	Send(string) error
	SetRecepient(string) error
}
