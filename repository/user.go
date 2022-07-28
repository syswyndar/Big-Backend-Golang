package repository

import (
	"Big-Backend-Golang/models"

	"gorm.io/gorm"
)

// type UserRepository interface {
// 	FindAll(db *gorm.DB) ([]models.User, error)
// 	FindById(Id int, db *gorm.DB) (models.User, error)
// 	CreateUser(user models.User, db *gorm.DB) (models.User, error)
// 	DeleteById(Id int, db *gorm.DB) (models.User, error)
// }

func CreateUser(user models.User, db *gorm.DB) (models.User, error) {
	err := db.Create(&user).Error

	return user, err
}
