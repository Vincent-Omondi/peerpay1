// database/migrations.go
package database

import (
	"log"

	"peerpay/models"
)

func Migrate() {
	err := DB.AutoMigrate(&models.User{}, &models.Share{}, &models.Admin{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
}
