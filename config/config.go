package config

import (
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

var ConfigData *Config = nil

// LoadConfig загружает и расшифровывает файл конфигурации YAML в структуру Config.
func LoadConfig(log *slog.Logger) (*Config, error) {
	// Расшифровываем файл
	fileData, err := os.ReadFile("local.yaml")
	if err != nil {
		log.Error("Error decrypting YAML file:", err)
		return nil, err
	}

	// Инициализируем структуру для конфигурации
	var conf Config

	// Распаковываем YAML данные в структуру Config
	err = yaml.Unmarshal(fileData, &conf)
	if err != nil {
		log.Error("Error unmarshalling YAML file:", err)
		return nil, err
	}

	return &conf, nil
}
