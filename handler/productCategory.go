package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	email := c.MustGet("Email")
	id := c.MustGet("ID")

	fmt.Println(email)
	fmt.Println(id)

	// var input request.CategoryRequest
	// db := c.MustGet("db").(*gorm.DB)

	// //binding data body to category request struct
	// err := c.ShouldBindJSON(&input)

	// // check if binding request error
	// if err != nil {
	// 	for _, err := range err.(validator.ValidationErrors) {
	// 		errorMessage := fmt.Sprintf("Error on field %s, condition %s", err.Field(), err.ActualTag())
	// 		c.JSON(http.StatusBadRequest, errorMessage)
	// 		return
	// 	}
	// }

	// // create category to database
	// models := models.Category_Product{
	// 	Category_name: input.Category_Name,
	// }

	// category, createCategoryError := repository.CreateCategory(models, db)

	// if createCategoryError != nil {
	// 	c.
	// }

}
