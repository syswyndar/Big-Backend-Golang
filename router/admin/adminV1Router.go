package adminRouter

import (
	"Big-Backend-Golang/handler"
	"Big-Backend-Golang/middleware"

	"github.com/gin-gonic/gin"
)

func AdminV1(r *gin.RouterGroup) {

	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)

	r.Use(middleware.Authentication)
	r.POST("/categories", middleware.AdminAuthorization, handler.CreateCategory)
	r.GET("/categories", handler.FindAllCategory)
	r.POST("/products", middleware.AdminAuthorization, middleware.UploadImage, handler.CreateProduct)
	r.GET("/products", handler.FindAllProduct)
}
