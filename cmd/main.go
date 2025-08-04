package main

import (
	"AutoBuckup/internal/config"
	"AutoBuckup/internal/log"
	"AutoBuckup/internal/services"
	"gopkg.in/yaml.v3"
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
	if readConfig.Debug {
		marshal, err := yaml.Marshal(readConfig)
		if err != nil {
			os.Exit(1)
		}
		log.Logger.Debug(string(marshal))
	}

	services.AutoBackup(readConfig)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan
	log.Logger.Infof("Received signal: %s exit.\n", sig)
}
