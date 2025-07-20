package main

import (
	"backend_golang/app"
	"backend_golang/config"
	"backend_golang/logger"
	"backend_golang/utils"
)

func main() {
	log := logger.InitLogger()
	config, err := config.LoadEnvConfig()
	if err != nil {
		log.WithField("error", err).Fatalf("unable to load config")
	}
	utils.RegisterCustomValidations()

	app := app.New(config, log)
	app.Run()
}
