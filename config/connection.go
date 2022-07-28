package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	// config database with driver postgres
	dsn := "host=localhost user=syswyndar password=postgres dbname=ecommerce-golang port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// chack if connection database error
	if err != nil {
		log.Fatal("DB Connection Error")
	}

	migration(db)

	return db
}
