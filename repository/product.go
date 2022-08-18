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

func FindProductById(product models.Product, db *gorm.DB, Id int) (models.Product, error) {
	err := db.Preload(clause.Associations).First(&product, Id).Error
	return product, err
}

func UpdateProduct(product models.Product, db *gorm.DB, id int) (models.Product, error) {
	err := db.Save(&product).Error
	return product, err
}

func DeleteProduct(product models.Product, db *gorm.DB, Id int) (models.Product, error) {
	err := db.Delete(&product, Id).Error
	return product, err
}
