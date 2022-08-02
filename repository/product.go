package repository

import (
	"Big-Backend-Golang/models"

	"gorm.io/gorm"
)

func CreateProduct(product models.Product, db *gorm.DB) (models.Product, error) {
	err := db.Create(&product).Error

	return product, err
}

func FindAllProduct(product models.Product, db *gorm.DB) ([]models.Product, error) {
	var productArr []models.Product
	err := db.Find(&productArr).Error
	return productArr, err
}
