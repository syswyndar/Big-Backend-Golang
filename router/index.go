package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MainRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	v1 := r.Group("/v1")
	routerV1(v1)

	return r
}
