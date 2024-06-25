package logger

import (
	"io"
	"os"

	"log/slog"
)

func SetupLogger() (*slog.Logger, error) { // Настройка логгера
	var log *slog.Logger
	file, err := os.OpenFile("./logs/log.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644) // создаем новый файл для логирования // |os.O_APPEND
	if err != nil {
		log.Error("error opening log file: %v", err)
		return nil, err
	}
	mw := io.MultiWriter(os.Stdout, file) // подключаем вывод в два потока: консоль и файл

	log = slog.New(slog.NewTextHandler(mw, &slog.HandlerOptions{Level: slog.LevelDebug}))

	return log, nil

}
