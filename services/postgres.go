package services

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"boilerplate/config"
	"boilerplate/models"
)

var Database *gorm.DB

func InitializeDatabase() {
	db, err := gorm.Open(postgres.Open(
		fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
			config.AppConfig.DBHost, config.AppConfig.DBPort, config.AppConfig.DBDatabase,
			config.AppConfig.DBUsername, config.AppConfig.DBPassword, config.AppConfig.SSLMode),
	), &gorm.Config{})

	if err != nil {
		panic("[ERROR] Failed to initialize database")
	}

	db.AutoMigrate(
		&models.Profile{},
		&models.Example{},
	)

	Database = db
}
