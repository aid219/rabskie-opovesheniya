package mailing

import "log/slog"

type Messager interface {
	Init(*slog.Logger) error
	Send(*slog.Logger, string, string, string, uint16) error
}
