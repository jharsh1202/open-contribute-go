package db

import (
	"log"

	"open-contribute/pkg/db/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init(dbPath string) *gorm.DB {
	// initialize a GORM database connection to a SQLite database
	// &gorm.Config{} provides default configuration options for GORM.
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Book{})

	return db
}
