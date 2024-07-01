package main

import (
	"rabiKrabi/config"
	"rabiKrabi/internal/logger"
	"rabiKrabi/internal/mailing/initial"
	"rabiKrabi/internal/rabbit"
)

var ConfigData config.Config

func main() {
	log, err := logger.SetupLogger()

	if err != nil {
		log.Error("Error setup logger", err)
	}
	config.ConfigData, err = config.LoadConfig(log, "local2.txt", []byte("qummy1234567890"))

	if err != nil {
		log.Error("Error load config", err)
	}

	ch, que, err := rabbit.Init(log, config.ConfigData.Rabbit.Host, config.ConfigData.Rabbit.Queue)

	if err != nil {
		log.Error("Error init rabbit", err)
	}

	senders := initial.InitAllSenders(log)

	err = rabbit.ReceiveAndSend(log, ch, que, senders)
	if err != nil {
		log.Error("Error receive and send operation", err)
	}

	for {
	}

}
