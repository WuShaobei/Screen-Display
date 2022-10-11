package router

import (
	"backEnd/service"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	userRouter(r)
	dataRouter(r)
}

func userRouter(r *gin.Engine) {
	g := r.Group("api/user")
	g.POST("/whoAmI", service.WhoAmIService)
	g.POST("/loginByPassword", service.LoginByPasswordService)
	g.POST("/loginBySessionKey", service.LoginBySessionKeyService)
	g.POST("/logout", service.LogoutService)
	g.POST("/register", service.RegisterService)
}

func dataRouter(r *gin.Engine) {
	g := r.Group("api/data")
	g.GET("/getBrandStatisticService", service.GetBrandStatisticService)
	g.POST("/postPaymentByYearService", service.PostPaymentByYearService)
	g.POST("/postOnlineOrderStatisticByYearAndMonthService", service.PostOnlineOrderStatisticByYearAndMonthService)
	g.POST("/postStatisticByYearService", service.PostStatisticByYearService)
	g.POST("/postFundingStatisticByYearAndMonthService", service.PostFundingStatisticByYearAndMonthService)

}
