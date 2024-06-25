package logger

import (
	"io"
	"os"
	"path/filepath"

	"log/slog"
)

func SetupLogger() (*slog.Logger, error) { // Настройка логгера
	var log *slog.Logger
	logDir := "./logs"
	logFile := filepath.Join(logDir, "log.txt")

	// Проверка и создание директории, если она не существует
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.MkdirAll(logDir, 0755)
		if err != nil {
			log.Error("Error to create log folder!")
			return nil, err
		}
	}

	// Открытие файла для логирования
	file, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Error("Error open log file!")
		return nil, err
	}

	// Настройка вывода в два потока: консоль и файл
	mw := io.MultiWriter(os.Stdout, file)

	// Создание логгера
	log = slog.New(slog.NewTextHandler(mw, &slog.HandlerOptions{Level: slog.LevelDebug}))

	return log, nil

}
