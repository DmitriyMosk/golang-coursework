package database

import (
	"fmt"
	"log"

	"golang-coursework/cmd/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Issue struct {
	gorm.Model
	ID          string `gorm:"primaryKey"`
	Key         string
	Summary     string
	Description string
}

type Project struct {
	ID   string `json:"id" gorm:"primaryKey"`
	Key  string `json:"key"`
	Name string `json:"name"`
}

type SearchResult struct {
	Issues []Issue `json:"issues"`
}

var DB *gorm.DB

func InitDB(cfg config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.DBSettings.DBHost, cfg.DBSettings.DBUser, cfg.DBSettings.DBPassword,
		cfg.DBSettings.DBName, cfg.DBSettings.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&Issue{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	return db, nil
}

func Create(db *gorm.DB, project *Project) error {
	return db.Create(project).Error
}
