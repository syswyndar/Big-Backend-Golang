package config

import (
	model "Big-Backend-Golang/models"

	"gorm.io/gorm"
)

func migration(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
		&model.User_Address{},
		&model.User_Profile{},
		&model.Category_Product{},
		&model.Product{},
		&model.Product_Image{},
		&model.Chart{},
		&model.Saved_Product{},
		&model.Transaction{},
		&model.History_Transaction{},
	)
}
