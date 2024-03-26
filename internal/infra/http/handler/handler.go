package handler

import (
	"time"

	"github.com/mbrunos/go-hire/config"
	"github.com/mbrunos/go-hire/internal/core/usecases"
	"github.com/mbrunos/go-hire/internal/infra/database/repository"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB, jwtSecret string, jwtExp time.Duration) (*JobHandler, *UserHandler) {

	jobRepo := repository.NewJobRepository(db)
	jobUseCase := usecases.NewJobUseCase(jobRepo)
	jobHandler := NewJobHandler(jobUseCase)

	userRepo := repository.NewUserRepository(db)
	userUseCase := usecases.NewUserUseCase(userRepo)
	userHandler := NewUserHandler(userUseCase, config.JWTSecret, config.JWTExp)

	return jobHandler, userHandler
}
