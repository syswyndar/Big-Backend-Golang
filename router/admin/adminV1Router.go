package adminRouter

import (
	"Big-Backend-Golang/handler"
	"Big-Backend-Golang/middleware"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminV1(r *gin.RouterGroup) {

	r.GET("/", func(c *gin.Context) {
		fmt.Println("server running")
		c.JSON(http.StatusOK, gin.H{
			"messsage": "Server Running",
		})
	})
	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)

	r.Use(middleware.Authentication)
	r.POST("/categories", middleware.AdminAuthorization, handler.CreateCategory)
	r.GET("/categories", handler.FindAllCategory)
	r.GET("/categories/:id", handler.FindCategoryById)
	r.POST("/products", middleware.AdminAuthorization, middleware.CreateProduct, handler.CreateProduct)
	r.GET("/products", handler.FindAllProduct)
	r.GET("/products/:id", handler.FindProductById)
	r.PUT("/products/:id", middleware.AdminAuthorization, middleware.UpdateProduct, handler.UpdateProduct)
	r.DELETE("/products/:id", middleware.AdminAuthorization, handler.DeleteProduct)
}
