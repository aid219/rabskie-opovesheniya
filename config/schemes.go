package config

type Telegram struct {
	Token string // Токен Telegram бота
}

// Mail представляет конфигурацию для отправки почты.
type Mail struct {
	Sender   string // Адрес отправителя
	Host     string // SMTP хост
	Port     int    // Порт SMTP сервера
	Username string // Имя пользователя для аутентификации SMTP
	Password string // Пароль для аутентификации SMTP
}

// Rabbit представляет конфигурацию для RabbitMQ.
type Rabbit struct {
	Host  string // Адрес хоста RabbitMQ
	Queue string // Имя очереди RabbitMQ
}

// Config представляет общую конфигурацию приложения, загружаемую из YAML файла.
type Config struct {
	Messangers []string `yaml:"messangers"` // Список мессенджеров (может содержать другие мессенджеры помимо Telegram)
	Telegram   Telegram `yaml:"telegram"`   // Конфигурация для Telegram бота
	Mail       Mail     `yaml:"email"`      // Конфигурация для отправки почты
	Rabbit     Rabbit   `yaml:"rabbit"`     // Конфигурация для RabbitMQ
}
