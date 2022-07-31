package repository

import (
	"Big-Backend-Golang/models"

	"gorm.io/gorm"
)

func CreateUser(user models.User, db *gorm.DB) (models.User, error) {
	err := db.Create(&user).Error

	return user, err
}

func GetUserByEmail(user models.User, email string, db *gorm.DB) (models.User, error) {
	err := db.First(&user, "email = ?", email).Error

	return user, err
}

func GetUserByID(user models.User, Id int, db *gorm.DB) (models.User, error) {
	err := db.First(&user, "ID = ?", Id).Error

	return user, err
}
