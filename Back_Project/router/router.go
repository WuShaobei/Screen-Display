package router

import (
	"Back_Project/service"
	users "Back_Project/service/ApiUsers"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	g := r.Group("api/v1")

	g.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "hello gin",
		})
	})

	g.POST("/loginByPassword", users.LoginByPassword)
	g.POST("/loginBySession", users.LoginBySession)
	g.POST("/create", users.CreateMember)
	g.POST("/whoAmI", users.WhoAmI)
	g.POST("/api1", service.Api1)
	g.POST("/api2", service.Api2)
	g.POST("/api3", service.Api3)
	g.POST("/api4", service.Api4)
	g.POST("/api5", service.Api5)
}
