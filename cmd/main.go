package main

import (
	"github.com/joho/godotenv"
	"github.com/mbrunos/go-hire/config"
	"github.com/mbrunos/go-hire/internal/infra/http/handler"
	"github.com/mbrunos/go-hire/internal/infra/http/routes"
	"github.com/mbrunos/go-hire/internal/infra/http/server"
	"github.com/mbrunos/go-hire/pkg/middleware"
)

func main() {
	logger := config.GetLogger()

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

	s := server.NewServer(config.ServerPort, logger)
	s.Router.Use(middleware.RequestLogger(logger))

	db := config.GetDB()
	jobHandler, userHandler := handler.Setup(db, config.JWTSecret, config.JWTExp)

	routes.Setup(s.Router, userHandler, jobHandler)

	if err := s.Start(); err != nil {
		logger.ErrorF("Error starting server: %s", err)
		return
	}
}
