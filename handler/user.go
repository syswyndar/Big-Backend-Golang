package handler

import (
	"Big-Backend-Golang/helpers"
	"Big-Backend-Golang/models"
	"Big-Backend-Golang/repository"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {
	var input models.User
	db := c.MustGet("db").(*gorm.DB)
	//binding data body to model User struct
	err := c.ShouldBindJSON(&input)

	//check if binding returning error
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition %s", err.Field(), err.ActualTag())
			c.JSON(http.StatusBadRequest, errorMessage)
			return
		}
	}

	//hashing password
	hashPassword, e := helpers.HashingPassword(input.Password)

	// check if error when hashing password
	if e != nil {
		c.JSON(http.StatusBadRequest, e)
		return
	}

	// change password from input user with hashed password
	input.Password = hashPassword

	// create data to database
	user, createUserError := repository.CreateUser(input, db)

	// check if create data to database error
	if createUserError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": createUserError,
		})
		return
	}

	// create custom response because  models.User have sensitive data
	result := models.User{
		ID:         user.ID,
		Email:      user.Email,
		Created_at: user.Created_at,
		Updated_at: user.Created_at,
		Deleted_at: user.Deleted_at,
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User Successfully registered",
		"data":    result,
	})
}

// func Login (c *gin.Context) {
// 	var input request.RegisterRequest

// }
