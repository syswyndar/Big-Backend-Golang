package middleware

import (
	"Big-Backend-Golang/helpers"
	"Big-Backend-Golang/models"
	"Big-Backend-Golang/repository"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Authentication(c *gin.Context) {
	// check req authorization from header
	checkToken := c.GetHeader("authorization")
	if len(checkToken) == 0 {
		message := "authorization header is not provided"
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": message,
		})
		return
	}

	// check type of request authorization header
	token := strings.Split(c.Request.Header["Authorization"][0], " ")
	if token[0] != "Bearer" {
		message := "wrong type of authorization header"
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": message,
		})
		return
	}

	// get data from token
	idUser := helpers.DecodeToken(token[1])

	db := c.MustGet("db").(*gorm.DB)
	user, err := repository.GetUserByID(models.User{}, idUser, db)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err,
		})
		return
	}

	c.Set("ID", user.ID)
	c.Set("Email", user.Email)
	c.Next()
}
