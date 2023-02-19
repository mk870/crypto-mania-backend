package db

import (
	"cryptoApi/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://lyvfmjek:1he2gHJkc7eIKS8MgbUTwtS_uee5tazd@trumpet.db.elephantsql.com/lyvfmjek"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{}, &models.Coin{}, &models.VerificationToken{})
	DB = db
}
