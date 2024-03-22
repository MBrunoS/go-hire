package main

import (
	"github.com/joho/godotenv"
	"github.com/mbrunos/go-hire/config"
	"github.com/mbrunos/go-hire/internal/core/usecases"
	"github.com/mbrunos/go-hire/internal/infra/database/repository"
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

	db := config.GetDB()

	jobRepo := repository.NewJobRepository(db)
	jobUseCase := usecases.NewJobUseCase(jobRepo)
	jobHandler := handler.NewJobHandler(jobUseCase)

	userRepo := repository.NewUserRepository(db)
	userUseCase := usecases.NewUserUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUseCase, config.JWTSecret)

	s := server.NewServer(config.ServerPort)
	s.Router.Use(middleware.Logger)

	routes.AddPublicRoutes(s.Router, userHandler, jobHandler)
	routes.AddPrivateRoutes(s.Router, userHandler, jobHandler)
	routes.AddSwaggerRoutes(s.Router)

	if err := s.Start(); err != nil {
		logger.ErrorF("Error starting server: %s", err)
		return
	}
}
