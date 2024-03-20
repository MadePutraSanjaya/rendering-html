package initializers

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() (*gorm.DB, error) {
	if DB != nil {
		return DB, nil
	}

	var err error
	DB, err = gorm.Open(sqlite.Open("delivery.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database:", err)
		return nil, err
	}

	return DB, nil
}
