package router

import (
	"Back_Project/service"
	users "Back_Project/service/ApiUsers"

	"github.com/gin-gonic/gin"
)

// RegisterRouter
//
//	@Description: 路由设置
//	@param r
//	@data 2022-09-28 13:55:22
//	@author WuShaobei
func RegisterRouter(r *gin.Engine) {
	g := r.Group("api/v1")

	g.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "hello gin",
		})
	})

	g.POST("/loginByPassword", users.LoginByPassword)
	g.POST("/loginBySession", users.LoginBySession)
	g.POST("/create", users.Register)
	g.POST("/whoAmI", users.WhoAmI)
	g.POST("/select/ChineseCateringStatistic", service.SelectChineseCateringStatistic)
	g.POST("/select/ChineseCateringOnlineOrderStatistic", service.SelectChineseCateringOnlineOrderStatistic)
	g.POST("/select/ChineseCateringPayment", service.SelectChineseCateringPayment)
	g.POST("/select/ChineseCateringBrandStatistic", service.SelectChineseCateringBrandStatistic)
	g.POST("/select/ChineseCateringFundingStatistic", service.SelectChineseCateringFundingStatistic)
}
