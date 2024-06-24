package mailing

type Messager interface {
	Init() error
	Send(string, string, string) error
}
