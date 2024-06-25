package config

type Telegram struct {
	Token string
}

type Mail struct {
	Sender   string
	Host     string
	Port     int
	Username string
	Password string
}

type Rabbit struct {
	Host  string
	Queue string
}

type Config struct {
	Messangers []string `yaml:"messangers"`
	Telegram   Telegram `yaml:"telegram"`
	Mail       Mail     `yaml:"mail"`
	Rabbit     Rabbit   `yaml:"rabbit"`
}
