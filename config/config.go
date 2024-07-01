package config

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
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

	// Проверяем, что длина ciphertext кратна размеру блока AES
	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	// Расшифровываем данные
	plaintext := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(block, key[:aes.BlockSize])
	mode.CryptBlocks(plaintext, ciphertext)

	// Удаляем дополнительные байты (padding)
	plaintext = PKCS5Unpad(plaintext)

	return plaintext, nil
}

// PKCS5Unpad удаляет дополнительные байты (padding) для CBC
func PKCS5Unpad(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}

// LoadConfig загружает и расшифровывает файл конфигурации YAML в структуру Config.
func LoadConfig(log *slog.Logger, encryptedFilename string, key []byte) (*Config, error) {
	// Расшифровываем файл
	decryptedData, err := decryptFile(encryptedFilename, key)
	if err != nil {
		fmt.Println("Error decrypting YAML file:", err)
		return nil, err
	}

	// Инициализируем структуру для конфигурации
	var conf Config

	// Распаковываем YAML данные в структуру Config
	err = yaml.Unmarshal(decryptedData, &conf)
	if err != nil {
		fmt.Println("Error unmarshalling YAML file:", err)
		return nil, err
	}

	return &conf, nil
}
