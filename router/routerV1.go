package router

import (
	"Big-Backend-Golang/handler"

	"github.com/gin-gonic/gin"
)

func routerV1(r *gin.RouterGroup) {
	r.POST("/register", handler.Register)
}
