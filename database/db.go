package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"blog/model"
)

var DB *gorm.DB

func Connect() {
	// DSN (Data Source Name)
	dsn := "host=localhost user=bloguser password=password dbname=blogdb port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	fmt.Println("Connected to the PostgreSQL database!")

	err = db.AutoMigrate(&model.Blog{})
	if err != nil {
		log.Fatal("Failed to migrate the database:", err)
	}

	fmt.Println("Blog table migrated successfully!")
	DB = db
}
