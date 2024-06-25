package config

import (
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadConfig(log *slog.Logger) *Config {
	var conf Config
	yamlFile, err := os.ReadFile("./local.yaml")
	if err != nil {
		log.Error("Error read yaml file", err)
		return nil
	}

	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Error("Error unmarshal yaml file", err)
		return nil
	}

	return &conf
}
