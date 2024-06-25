package main

import (
	"rabiKrabi/config"
	"rabiKrabi/globals"
	"rabiKrabi/internal/logger"
	"rabiKrabi/internal/mailing"
	"rabiKrabi/internal/mailing/initial"
	"rabiKrabi/internal/rabbit"
)

var ConfigData config.Config

// 1150762777 tg
func main() {
	log, err := logger.SetupLogger()

	if err != nil {
		log.Error("Error setup logger", err)
	}

	globals.ConfigData = config.LoadConfig(log)

	ch, que, err := rabbit.Init(log, globals.ConfigData.Rabbit.Host, globals.ConfigData.Rabbit.Queue)

	if err != nil {
		log.Error("Error init rabbit", err)
	}

	senders := initial.InitAllSenders(log)

	incomeCh, _ := rabbit.Receive(log, ch, que)

	for {
		incomeData := <-incomeCh
		if len(incomeData.To) > 0 {
			go mailing.Send(log, incomeData, senders)
		}

	}
}
