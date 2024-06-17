package db

import (
	"github.com/blanc42/mooca/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	// conenct to postgresdb
	dsn := "host=localhost user=postgres password=password dbname=mooca port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db
}

func InitializeDB() {
	DB.AutoMigrate(&models.User{}, &models.Role{})
}
