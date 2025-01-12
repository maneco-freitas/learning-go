package database

import (
	"api/internal/config"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDBConnection(cfg *config.Config) *gorm.DB {
	var db *gorm.DB
	var err error

	switch cfg.Database.Type {
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(cfg.Database.Path), &gorm.Config{})
	default:
		log.Fatalf("Unsupported database type: %s", cfg.Database.Type)
	}
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}
