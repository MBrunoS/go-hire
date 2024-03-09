package config

import (
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	db, err = InitDB()
	if err != nil {
		return err
	}

	logger = NewLogger()
	return nil
}

func GetDB() *gorm.DB {
	return db
}

func GetLogger() *Logger {
	return logger
}
