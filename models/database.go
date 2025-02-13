package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDatabase() (*gorm.DB, error) {
	var err error
	dsn := "host=localhost user=user password=password dbname=spy_cat_agency port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	DB.AutoMigrate(&Cat{}, &Mission{}, &Target{})
	return DB, nil
}
