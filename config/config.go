package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Messangers []string `yaml:"messangers"`
}

func LoadConfig() *Config {
	var conf Config
	yamlFile, err := os.ReadFile("./local.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return &conf
}
