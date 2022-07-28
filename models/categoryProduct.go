package models

import (
	"time"
)

type Category_Product struct {
	ID            int    `gorm:"primary_key;AUTO_INCREMENT"`
	Category_name string `gorm:"type:varchar(30);"`
	Created_at    time.Time
	Updated_at    time.Time
	Deleted_at    *time.Time
}
