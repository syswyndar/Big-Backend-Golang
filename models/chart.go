package models

import (
	"time"
)

type Chart struct {
	ID         int     `gorm:"primary_key;AUTO_INCREMENT"`
	User_Id    int     `gorm:"not null;index"`
	User       User    `gorm:"not null;foreignkey:User_Id"`
	Product_Id int     `gorm:"not null;index"`
	Product    Product `gorm:"foreignkey:Product_Id"`
	Created_at time.Time
	Updated_at time.Time
	Deleted_at *time.Time
}
