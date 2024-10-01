package database

import (
	"log"
	"trilune/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Connect to database SQLite
	var err error
	DB, err = gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	} else {
		log.Println("Database connected successfully!")
	}

	// Автоматическая миграция модели Cocktail
	DB.AutoMigrate(&models.Expenses{}, &models.Category{}, &models.Incomes{})

}
