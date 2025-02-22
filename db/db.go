package db

import (
	"fmt"
	"log"

	"example.com/short_link/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{}) // Assigning to the global DB

	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	fmt.Println("🚀 Database connection established")

	// Migrate the schema
	err = DB.AutoMigrate(&models.ShortURL{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	fmt.Println("🚀 Migration done")
}
