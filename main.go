package main

import (
	"rabiKrabi/config"
	"rabiKrabi/internal/logger"
	"rabiKrabi/internal/mailing"
	"rabiKrabi/internal/mailing/initial"
	"rabiKrabi/internal/rabbit"
	"time"
)

var ConfigData config.Config

func main() {
	time.Sleep(time.Second * 10)
	log, err := logger.SetupLogger()

	if err != nil {
		log.Error("Error setup logger", err)
	}

	config.ConfigData = config.LoadConfig(log)

	ch, que, err := rabbit.Init(log, config.ConfigData.Rabbit.Host, config.ConfigData.Rabbit.Queue)

	if err != nil {
		log.Error("Error init rabbit", err)
	}

	senders := initial.InitAllSenders(log)

	err = mailing.Receive(log, ch, que, senders)
	if err != nil {
		log.Error("Error receive", err)
	}
	select {}
}
