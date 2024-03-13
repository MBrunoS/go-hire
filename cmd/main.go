package main

import (
	"github.com/joho/godotenv"
	"github.com/mbrunos/go-hire/config"
	"github.com/mbrunos/go-hire/router"
)

var logger *config.Logger

func main() {
	logger = config.GetLogger()

	err := godotenv.Load()

	if err != nil {
		logger.ErrorF("Error loading .env file: %s", err)
		return
	}

	err = config.Init()
	if err != nil {
		logger.ErrorF("Error initializing config: %s", err)
		return
	}

	router.Setup()
}
