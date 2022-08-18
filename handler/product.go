package handler

import (
	"Big-Backend-Golang/models"
	"Big-Backend-Golang/repository"
	"Big-Backend-Golang/request"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateProduct(c *gin.Context) {
	price, _ := strconv.ParseInt(c.PostForm("productPrice"), 10, 64)
	stock, _ := strconv.Atoi(c.PostForm("productStock"))
	categoryId, _ := strconv.Atoi(c.PostForm("productCategoryId"))

	input := &request.ProductRequest{
		Product_Name:        c.PostForm("productName"),
		Product_Description: c.PostForm("productDescription"),
		Product_Price:       price,
		Product_Stock:       stock,
		Product_Category_Id: categoryId,
	}

	db := c.MustGet("db").(*gorm.DB)
	product_img := c.MustGet("product_img").(string)

	// create product unique id
	unixtime := strconv.FormatInt(time.Now().Unix(), 10)
	currentTime := time.Now()
	currentime := strings.Join(strings.Split(currentTime.Format("2006-01-02"), "-"), "")

	// binding to model Product
	product := models.Product{
		Product_Unique_Id:   currentime + "_" + unixtime,
		Product_Name:        input.Product_Name,
		Product_Description: input.Product_Description,
		Product_Price:       input.Product_Price,
		Product_Stock:       input.Product_Stock,
		Category_Product_Id: input.Product_Category_Id,
		Product_Main_Img:    product_img,
	}

	// create data to database
	createdProduct, createErr := repository.CreateProduct(product, db)

	// check if create product error
	if createErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": createErr,
		})
		return
	}

	// sent response to frontend if product successfully created
	c.JSON(http.StatusCreated, gin.H{
		"message": "product successfully created",
		"data":    createdProduct,
	})
}

func FindAllProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	model := models.Product{}

	// find all product in database
	product, findingError := repository.FindAllProduct(model, db)

	// check if finding product error
	if findingError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": findingError,
		})
		return
	}

	// send response to frontend if finding product success
	c.JSON(http.StatusOK, gin.H{
		"message": "get all product success",
		"data":    product,
	})
}

func FindProductById(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	Id, err := strconv.Atoi(c.Param("id")) // get id param and convert type data from string to int

	// check if user doesn't send id in req.param
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id param is required",
		})
	}

	// find product in database
	product, findingError := repository.FindProductById(models.Product{}, db, Id)

	// check if finding data in database error
	if findingError != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Product with Id: " + strconv.Itoa(Id) + " is not found",
		})
		return
	}

	// send response data to frontend if success find product
	c.JSON(http.StatusOK, gin.H{
		"message": "find product by id Success",
		"data":    product,
	})
}

func UpdateProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	Id, err := strconv.Atoi(c.Param("id")) // get id param and convert type data from string to int
	product_img := c.MustGet("product_img").(string)
	fmt.Println(product_img, "ini product image")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id param is required",
		})
		return
	}

	price, _ := strconv.ParseInt(c.PostForm("productPrice"), 10, 64)
	stock, _ := strconv.Atoi(c.PostForm("productStock"))
	categoryId, _ := strconv.Atoi(c.PostForm("productCategoryId"))

	// find product in database
	finding, findingError := repository.FindProductById(models.Product{}, db, Id)

	// check if finding data in database error
	if findingError != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Product with Id: " + strconv.Itoa(Id) + " is not found",
		})
		return
	}

	product := models.Product{
		ID:                  finding.ID,
		Product_Name:        c.PostForm("productName"),
		Product_Description: c.PostForm("productDescription"),
		Product_Unique_Id:   finding.Product_Unique_Id,
		Product_Price:       price,
		Product_Stock:       stock,
		Category_Product_Id: categoryId,
		Product_Main_Img:    product_img,
	}

	updateProduct, errUpdateProduct := repository.UpdateProduct(product, db, Id)

	if errUpdateProduct != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errUpdateProduct,
		})
		return
	}

	// send response data to frontend if success find product
	c.JSON(http.StatusOK, gin.H{
		"message": "update product by id Success",
		"data":    updateProduct,
	})
}

func DeleteProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	Id, err := strconv.Atoi(c.Param("id")) // get id param and convert type data from string to int

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id param is required",
		})
		return
	}

	_, errDeleteProduct := repository.DeleteProduct(models.Product{}, db, Id)

	if errDeleteProduct != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errDeleteProduct,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "product with id" + strconv.Itoa(Id) + " success deleted",
	})
}
