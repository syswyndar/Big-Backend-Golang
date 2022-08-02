package repository

import (
	"Big-Backend-Golang/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateProduct(product models.Product, db *gorm.DB) (models.Product, error) {
	err := db.Create(&product).Error

	return product, err
}

func FindAllProduct(product models.Product, db *gorm.DB) ([]models.Product, error) {
	var productArr []models.Product
	err := db.Preload(clause.Associations).Find(&productArr).Error

	return productArr, err
}
