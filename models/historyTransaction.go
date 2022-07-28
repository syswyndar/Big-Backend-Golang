package models

import (
	"time"
)

type History_Transaction struct {
	// gorm.Model
	ID                  int    `gorm:"primary_key;AUTO_INCREMENT"`
	User_Id             int    `gorm:"not null;index"`
	User                User   `gorm:"not null;foreignkey:User_Id"`
	Product_Unique_Id   string `gorm:"type:varchar(20);"`
	Product_Name        string `gorm:"type:varchar(30);"`
	Product_Description string `gorm:"type:text"`
	Product_Price       int    `gorm:"type:int;"`
	Product_Category    string `gorm:"type:varchar(30);"`
	Quantity            int    `gorm:"type:int;"`
	Total_Price         int    `gorm:"type:int;"`
	Created_at          time.Time
	Updated_at          time.Time
	Deleted_at          *time.Time
}
