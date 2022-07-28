package main

import (
	"Big-Backend-Golang/config"
	"Big-Backend-Golang/router"
	"log"

	// "Big-Backend-Golang/config"
	// "Big-Backend-Golang/router"

	// "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("can't load file .env")
	}

	// call connection database
	connection := config.Connection()
	r := router.MainRouter(connection)

	r.Run()
}
