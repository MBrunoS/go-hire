package handler

import (
	"github.com/mbrunos/go-hire/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func Init() {
	logger = config.GetLogger()
	db = config.GetDB()
}
