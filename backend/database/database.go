// database/database.go
package database

import (
	"fmt"
	"log"
	"os"

	"peerpay/backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		getEnv("DB_USER", "sql5712813"),
		getEnv("DB_PASSWORD", "ZJavBXXU7s"),
		getEnv("DB_HOST", "sql5.freesqldatabase.com"),
		getEnv("DB_PORT", "3306"),
		getEnv("DB_NAME", "sql5712813"),
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	} else {
		log.Println("Database connection established")
	}

	models.DB = DB // Assign the connection to models.DB
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
