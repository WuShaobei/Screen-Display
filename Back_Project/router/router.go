package router

import (
	"server_side/function"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	g := r.Group("api/v1")

	g.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "hello gin",
		})
	})

	g.POST("/insert", function.ApiInsert)
	g.POST("/select", function.ApiSelect)
}
