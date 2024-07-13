package db

import (
	"os"

	"github.com/blanc42/mooca/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	// Connect to PostgreSQL database using DSN from environment variable
	dsn := getDSNFromEnv()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db
}

func getDSNFromEnv() string {
	dsn := os.Getenv("DSN")

	if dsn == "" {
		panic("DSN environment variable is not set")
	}
	return dsn
}

func InitializeDB() {
	DB.AutoMigrate(&models.User{}, &models.Role{}, &models.Institute{}, &models.Teacher{}, &models.Course{})
}
