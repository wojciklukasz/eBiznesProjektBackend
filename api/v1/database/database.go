package database

import (
	"ProjektBackend/api/v1/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB = nil

func Connect() {
	db, err := gorm.Open(sqlite.Open("store.db"))
	if err != nil {
		panic("DATABASE ERROR")
	}

	err = db.AutoMigrate(&models.User{})
	err = db.AutoMigrate(&models.Category{})
	err = db.AutoMigrate(&models.Manufacturer{})
	err = db.AutoMigrate(&models.Product{})
	err = db.AutoMigrate(&models.Order{})
	if err != nil {
		panic("MIGRATION FAILED")
	}

	Database = db
}
