package router

import (
	"Back_Project/manager"
	"Back_Project/service/users"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	g := r.Group("api/v1")

	g.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "hello gin",
		})
	})

	g.POST("/login", users.Login)
}

func RegisterRouterTest(r *gin.Engine) {
	g := r.Group("api/test")

	g.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "hello gin",
		})
	})

	g.GET("/select/first", manager.TestApiSelectFirst)
}
