package mysql

import (
	"belajar-golang/internal/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	// Initialize mysql connection
	dsn := "root:@tcp(127.0.0.1:3306)/belajar_golang?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto migrate
	db.AutoMigrate(&entity.User{})

	return db, nil
}
