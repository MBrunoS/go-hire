package config

import (
	"os"
	"time"

	"gorm.io/gorm"
)

var (
	db         *gorm.DB
	logger     *Logger
	ServerPort string
	JWTSecret  string
	JWTExp     time.Duration
)

func Init() error {
	var err error

	db, err = InitDB()
	if err != nil {
		return err
	}

	ServerPort = os.Getenv("SERVER_PORT")
	JWTSecret = os.Getenv("JWT_SECRET")
	JWTExp, err = time.ParseDuration(os.Getenv("JWT_EXP"))

	if err != nil {
		return err
	}

	return nil
}

func GetDB() *gorm.DB {
	return db
}

func GetLogger() *Logger {
	if logger == nil {
		logger = NewLogger()
	}
	return logger
}
