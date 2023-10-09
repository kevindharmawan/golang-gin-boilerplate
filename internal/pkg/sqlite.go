package pkg

import (
	"boilerplate/internal/features/example"
	"boilerplate/internal/features/user"
	"boilerplate/internal/shared/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSqlite(dbConfig config.DatabaseConfig) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbConfig.FileName), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(example.ExampleGorm{}, &user.UserGorm{})

	return db, nil
}
