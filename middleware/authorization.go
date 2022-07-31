package middleware

import (
	"Big-Backend-Golang/models"
	"Big-Backend-Golang/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Config struct {
	Id int
}

func AdminAuthorization(c *gin.Context) {
	id := c.MustGet("ID").(int)
	db := c.MustGet("db").(*gorm.DB)

	// // check user in database
	user, err := repository.GetUserByID(models.User{}, id, db)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	if user.Role != "Admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "your role is restricted to access this endpoint",
		})
		return
	}

	c.Next()
}
