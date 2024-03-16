package config

import (
	"os"

	"gorm.io/gorm"
)

var (
	db         *gorm.DB
	logger     *Logger
	ServerPort string
)

func Init() error {
	var err error

	db, err = InitDB()
	if err != nil {
		return err
	}

	ServerPort = os.Getenv("SERVER_PORT")

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
