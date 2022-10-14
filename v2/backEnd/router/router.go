package router

import (
	"backEnd/service"

	"github.com/gin-gonic/gin"
)

// RegisterRouter
//
//	@Description: URL管理
//	@param r
//	@data 2022-10-13 21:44:48
//	@author WuShaobei
func RegisterRouter(r *gin.Engine) {
	userRouter(r)
	dataRouter(r)
}

// userRouter
//
//	@Description: 用户类 URL
//	@param r
//	@data 2022-10-13 21:48:04
//	@author WuShaobei
func userRouter(r *gin.Engine) {
	g := r.Group("api/user")
	g.POST("/whoAmI", service.WhoAmIService)
	g.POST("/loginByPassword", service.LoginByPasswordService)
	g.POST("/loginBySessionKey", service.LoginBySessionKeyService)
	g.POST("/logout", service.LogoutService)
	g.POST("/register", service.RegisterService)
}

// dataRouter
//
//	@Description: 数据类 URL
//	@param r
//	@data 2022-10-13 21:48:20
//	@author WuShaobei
func dataRouter(r *gin.Engine) {
	g := r.Group("api/data")
	g.GET("/getAllDataFromBrandStatistic", service.GetAllDataFromBrandStatisticService)
	g.POST("/postAmountAndPercentageByYearFromStatistic", service.PostAmountAndPercentageByYearFromStatisticService)
	g.POST("/postAvgSalaryAndCountsByYearFromPayment", service.PostAvgSalaryAndCountsByYearFromPaymentService)
	g.POST("/postAmountByYearAndMonthFromOnlineOrderStatistic", service.PostAmountByYearAndMonthFromOnlineOrderStatisticService)
	g.GET("/getAllDataByFromFundingStatistic", service.GetAllDataByFromFundingStatisticService)
}
