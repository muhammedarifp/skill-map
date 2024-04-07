package database

import (
	"github.com/muhammedarifp/skill-map/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitDatabase() error {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return err
	}

	db.AutoMigrate(&models.UserModel{})
	return nil
}

func GetDb() *gorm.DB {
	return db
}
