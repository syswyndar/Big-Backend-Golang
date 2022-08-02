package repository

import (
	"Big-Backend-Golang/models"

	"gorm.io/gorm"
)

func CreateCategory(category models.Category_Product, db *gorm.DB) (models.Category_Product, error) {
	err := db.Create(&category).Error

	return category, err
}

func FindAllCategory(category models.Category_Product, db *gorm.DB) ([]models.Category_Product, error) {
	var categoryArr []models.Category_Product
	err := db.Find(&categoryArr).Error

	return categoryArr, err
}

func FindCategoryById(category models.Category_Product, db *gorm.DB, Id int) (models.Category_Product, error) {
	err := db.First(&category, Id).Error
	return category, err
}
