package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         int    `gorm:"primaryKey;uniqueIndex"`
	Email      string `gorm:"unique;not null;" json:"email" binding:"required"`
	Password   string `gorm:"not null;" json:"password" binding:"required"`
	Role       string
	Created_at time.Time
	Updated_at time.Time
	Deleted_at gorm.DeletedAt
}
