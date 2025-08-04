package main

import (
	"AutoBuckupG/internal/config"
	"AutoBuckupG/internal/log"
	"AutoBuckupG/internal/services"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Init()
	config.Init()
	readConfig := config.ReadConfig()
	if readConfig == nil {
		log.Logger.Fatal("Can't read config")
	}

	services.AutoBackup(readConfig)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan
	log.Logger.Infof("Received signal: %s exit.\n", sig)
}
