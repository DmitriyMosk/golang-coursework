package main

import (
	"flag"
	"golang-coursework/backend/resource/config"
	"golang-coursework/backend/resource/internal/app"
	"golang-coursework/backend/resource/pkg/logger"
)

func main() {
	configPath := flag.String("configPath", "backend/resource/config/config-resource.yaml", "Path to the config file")
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
			log.Log(logger.ERROR, err.Error())
			panic(err)
		}
	}(newApp)

	if err = newApp.Run(); err != nil {
		log.Log(logger.ERROR, err.Error())
		panic(err)
	}
}
