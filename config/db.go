package config

import (
	"os"

	"github.com/mbrunos/go-hire/internal/infra/database/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	pass := os.Getenv("MYSQL_ROOT_PASSWORD")
	dbname := os.Getenv("MYSQL_DATABASE")
	logger := GetLogger()

	dsn := "root:" + pass + "@tcp(127.0.0.1:3306)/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.ErrorF("Error connecting to database: %s", err)
		return nil, err
	}

	err = db.AutoMigrate(&model.Job{}, &model.User{})
	if err != nil {
		logger.ErrorF("Error migrating schema: %s", err)
		return nil, err
	}

	return db, nil
}
