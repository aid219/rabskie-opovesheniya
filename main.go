package main

import (
	"os"
	"os/signal"
	"rabiKrabi/config"
	"rabiKrabi/internal/logger"
	"rabiKrabi/internal/mailing"
	"rabiKrabi/internal/mailing/initial"
	"rabiKrabi/internal/rabbit"
	"syscall"
)

var ConfigData config.Config

func main() {
	log, err := logger.SetupLogger()

	if err != nil {
		log.Error("Error setup logger, critical stop", err)
		os.Exit(1)
	}
	config.ConfigData, err = config.LoadConfig(log)

	if err != nil {
		log.Error("Error load config", err)
	}

	ch, que, err := rabbit.Init(log, config.ConfigData.Rabbit.Host, config.ConfigData.Rabbit.Queue)

	if err != nil {
		log.Error("Error init rabbit, critical stop", err)
		os.Exit(1)
	}

	senders, err := initial.InitAllSenders(log)

	if err != nil {
		log.Error("Error init senders, critical stop", err)
		os.Exit(1)
	}
	err = mailing.Receive(log, ch, que, senders)
	if err != nil {
		log.Error("Error receive", err)
	}

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	sig := <-stop

	log.Info("Received signal: %s", sig)

	log.Info("Shutting down gracefully...")
}
