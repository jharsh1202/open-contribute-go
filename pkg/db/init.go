package db

import (
	"log"
	"open-contribute/pkg/db/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Init initializes the database and returns the database handler.
func Init(dbPath string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	err = db.AutoMigrate(&models.Book{})
	if err != nil {
		return nil, err
	}

	log.Printf("Database connected successfully to %s", dbPath)
	return db, nil
}
