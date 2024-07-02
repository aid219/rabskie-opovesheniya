package mailing

import "log/slog"

// Messager - интерфейс определяет методы, которые должен реализовать любой мессенджер (например, Telegram или Email).
type Messager interface {
	// Init инициализирует мессенджер
	Init(*slog.Logger) error

	// Send отправляет сообщение через мессенджер.
	// Принимает log для логирования, chatID/получателя, тему сообщения, содержание сообщения и задержку перед отправкой.
	Send(*slog.Logger, string, string, string, uint16) error
}
