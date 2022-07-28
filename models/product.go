package models

import (
	"time"
)

type Product struct {
	ID                  int              `gorm:"primary_key;AUTO_INCREMENT"`
	Product_Unique_Id   string           `gorm:"type:varchar(30);"`
	Product_Name        string           `gorm:"type:varchar(30);"`
	Product_Description string           `gorm:"type:text"`
	Product_Price       int              `gorm:"type:int;"`
	Product_Main_Img    string           `gorm:"type:text"`
	Product_Stock       int              `gorm:"type:int;"`
	Product_Category_Id int              `gorm:"not null;aindex"`
	Product_Category    Category_Product `gorm:"not null;foreignkey:Product_Category_Id"`
	Created_at          time.Time
	Updated_at          time.Time
	Deleted_at          *time.Time
}
