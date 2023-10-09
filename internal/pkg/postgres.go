package pkg

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"boilerplate/internal/shared/config"
)

func InitializePostgres(dbConfig config.DatabaseConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(
		fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
			dbConfig.Host, dbConfig.Port, dbConfig.Name, dbConfig.User, dbConfig.Password, dbConfig.SSLMode),
	), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
