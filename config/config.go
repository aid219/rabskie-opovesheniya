package config

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"io/ioutil"
	"log/slog"

	"gopkg.in/yaml.v3"
)

var ConfigData *Config = nil

func decryptFile(filename string, key []byte) ([]byte, error) {
	ciphertext, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

// LoadConfig loads and decrypts a YAML configuration file into a Config struct.
func LoadConfig(log *slog.Logger, encryptedFilename string, key []byte) (*Config, error) {
	// Расшифровываем файл
	decryptedData, err := decryptFile(encryptedFilename, key)
	if err != nil {
		log.Error("Error decrypting YAML file", err)
		return nil, err
	}

	// Инициализируем структуру для конфигурации
	var conf Config

	// Распаковываем YAML данные в структуру Config
	err = yaml.Unmarshal(decryptedData, &conf)
	if err != nil {
		log.Error("Error unmarshalling YAML file", err)
		return nil, err
	}

	return &conf, nil
}
