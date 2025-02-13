package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

var DB *gorm.DB

func SetupDatabase() (*gorm.DB, error) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	DB, err = gorm.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	DB.AutoMigrate(&Item{})
	return DB, nil
}

type Item struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Name  string `json:"name" binding:"required"`
	Price int    `json:"price" binding:"required"`
}
