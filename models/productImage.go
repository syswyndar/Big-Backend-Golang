package models

import (
	"time"
)

type Product_Image struct {
	// gorm.Model
	ID         int     `gorm:"primary_key;AUTO_INCREMENT"`
	Product_Id int     `gorm:"index"`
	Product    Product `gorm:"foreignkey:Product_Id"`
	Image      string  `gorm:"type:text"`
	Created_at time.Time
	Updated_at time.Time
	Deleted_at *time.Time
}
