package main

import (
	"github.com/joho/godotenv"
	"github.com/mbrunos/go-hire/config"
	"github.com/mbrunos/go-hire/internal/infra/http/handler"
	"github.com/mbrunos/go-hire/internal/infra/http/routes"
	"github.com/mbrunos/go-hire/internal/infra/http/server"
	"github.com/mbrunos/go-hire/pkg/middleware"
)

// @title Go Hire API
// @version 1
// @description This is a simple API for managing job offers and users.

// @contact.name Maurício Bruno da Silva
// @contact.url https://mbrunos.dev
// @contact.email contact@mbrunos.dev

// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization

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
