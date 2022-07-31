package router

import (
	"Big-Backend-Golang/handler"
	"Big-Backend-Golang/middleware"

	"github.com/gin-gonic/gin"
)

func routerV1(r *gin.RouterGroup) {
	// admin v1 routes
	admin := r.Group("/admin")
	{
		admin.POST("/register", handler.Register)
		admin.POST("/login", handler.Login)

		admin.Use(middleware.Authentication)
		admin.POST("/categories", middleware.AdminAuthorization, handler.CreateCategory)
	}
}
