package models

import (
	"time"
)

type User_Profile struct {
	ID            int    `gorm:"primary_key;AUTO_INCREMENT"`
	User          User   `gorm:"not null;foreignkey:User_Id"`
	User_Id       int    `gorm:"not null;index"`
	Fullname      string `gorm:"type:varchar(30);"`
	Date_Of_Birth string `gorm:"type:varchar(30);"`
	Gender        string `gorm:"type:varchar(30);"`
	Telephone     string `gorm:"type:varchar(30);"`
	User_Image    string `gorm:"type:text"`
	Created_at    time.Time
	Updated_at    time.Time
	Deleted_at    *time.Time
}
