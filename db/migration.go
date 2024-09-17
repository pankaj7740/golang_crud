package db

import (
	"fmt"
	"log"
	"crud_basic/model"
)

// AutoMigrate runs the auto-migration for the database models
func AutoMigrate() {
	err := DB.AutoMigrate(&model.Company{}, &model.User{})
	if err != nil {
		log.Fatalf("Failed to run auto-migration: %v", err)
	}
	fmt.Println("Database migration completed successfully")
}
