package handler

import (
	"github.com/mbrunos/go-hire/config"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() {
	db = config.GetDB()
}
