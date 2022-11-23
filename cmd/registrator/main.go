package main

import (
	"log"

	"github.com/matsuev/klsh-registrator/internal/app"
	"github.com/matsuev/klsh-registrator/internal/config"
)

func main() {
	log.Println("KLSH Registrator service")

	cfg, err := config.New()
	if err != nil {
		log.Fatalf("create config aborted with error: %v\n", err)
	}

	appInstance, err := app.New(cfg)
	if err != nil {
		log.Fatalf("create application aborted with error: %v\n", err)
	}

	appInstance.Run()
}
