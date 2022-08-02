package handler

import (
	"Big-Backend-Golang/models"
	"Big-Backend-Golang/repository"
	"Big-Backend-Golang/request"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func CreateCategory(c *gin.Context) {
	var input request.CategoryRequest

	// var input request.CategoryRequest
	db := c.MustGet("db").(*gorm.DB)

	//binding data body to category request struct
	err := c.ShouldBindJSON(&input)

	// check if binding request error
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition %s", err.Field(), err.ActualTag())
			c.JSON(http.StatusBadRequest, errorMessage)
			return
		}
	}

	// binding input data to product category models
	models := models.Category_Product{
		Category_name: input.Category_Name,
	}

	// create data to database
	category, createCategoryError := repository.CreateCategory(models, db)

	if createCategoryError != nil {
		c.JSON(http.StatusInternalServerError, createCategoryError)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Create Product Category success",
		"data":    category,
	})

}
