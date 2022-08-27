package config

import (
	"gin-gorm-rest/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=5432 "
	// db, err := gorm.Open(postgres.Open("postgres://gorm:gorm@localhost:5432/gorm"), &gorm.Config{})
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{})
	DB = db
}
