package main

import (
	"flag"
	"fmt"
	"golang-coursework/connector/config"
	"golang-coursework/connector/internal/app"
	"golang-coursework/connector/pkg/logger"
)

func main() {
	configPath := flag.String("configPath", "connector/config/config-connector.yaml", "Path to the config file")
	flag.Parse()

	log := logger.CreateNewLogger()

	cfg, err := config.NewConfig(*configPath)

	if err != nil {
		log.Log(logger.ERROR, err.Error())
		panic(err)
	}

	newApp, err := app.NewApp(cfg, log)

	if err != nil {
		log.Log(logger.ERROR, err.Error())
		panic(err)
	}

	defer func(newApp *app.App) {
		err := newApp.Close()
		if err != nil {
			fmt.Println(err.Error())

			log.Log(logger.ERROR, err.Error())
			panic(err)
		}
	}(newApp)

	if err = newApp.Run(); err != nil {
		log.Log(logger.ERROR, err.Error())
		panic(err)
	}
}
