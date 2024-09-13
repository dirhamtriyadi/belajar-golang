package sqlite

import (
	"belajar-golang/internal/entity"

	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// InitDB initializes the database
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&entity.User{})

	return db, nil
}
