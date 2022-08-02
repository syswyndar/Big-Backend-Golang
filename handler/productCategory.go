package handler

import (
	"Big-Backend-Golang/models"
	"Big-Backend-Golang/repository"
	"Big-Backend-Golang/request"
	"fmt"
	"net/http"
	"strconv"

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

func FindAllCategory(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	model := models.Category_Product{}

	// find all category in database
	categories, findingError := repository.FindAllCategory(model, db)

	if findingError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": findingError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "find all category success",
		"data":    categories,
	})
}

func FindCategoryById(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	Id, err := strconv.Atoi(c.Param("id"))

	// check if user doesn't send id param
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Id param is required",
		})
	}

	// find data category in database
	category, findingError := repository.FindCategoryById(models.Category_Product{}, db, Id)

	// check if error finding data
	if findingError != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Category with Id: " + strconv.Itoa(Id) + " is not found",
		})
		return
	}

	// send response to frontend if finding data success
	c.JSON(http.StatusOK, gin.H{
		"message": "find category by id success",
		"data":    category,
	})
}
