package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("Configuration loading error: %v", err)
	}
	dsn := fmt.Sprintf("host=%s dbname=%s password=%s user=%s port=%s sslmode=disable",
		cfg.DatabaseHost, cfg.DatabaseName, cfg.DatabasePassword, cfg.DatabaseUser, cfg.DatabasePort)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}

	fmt.Println("Successfully connected to the database")
	return DB
}
