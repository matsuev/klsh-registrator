package main

import (
	"log"

	"github.com/matsuev/klsh-registrator/internal/app"
	"github.com/matsuev/klsh-registrator/internal/config"
	"github.com/matsuev/klsh-registrator/internal/logging"
)

func main() {
	println("KLSH Registrator service")

	cfg, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}

	logger, err := logging.New(cfg.Logger)
	if err != nil {
		log.Fatalln(err)
	}

	appInstance, err := app.New(cfg, logger)
	if err != nil {
		logger.Fatalln(err)
	}

	if err := appInstance.Run(); err != nil {
		logger.Fatalln(err)
	}
}
