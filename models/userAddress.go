package models

import (
	"time"
)

type User_Address struct {
	ID              int    `gorm:"primaryKey;uniqueIndex"`
	User_Id         int    `gorm:"index;not null;"`
	User            User   `gorm:"not null;foreignkey:User_Id"`
	Province        string `gorm:"type:varchar(30);"`
	District        string `gorm:"type:varchar(30);"`
	Sub_District    string `gorm:"type:varchar(30);"`
	Full_Address    string `gorm:"type:varchar(100);"`
	Is_Main_Address bool   `gorm:"default:true"`
	Created_at      time.Time
	Updated_at      time.Time
	Deleted_at      *time.Time
}
