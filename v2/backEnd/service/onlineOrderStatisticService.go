// Package service
// @Description: 网上订单相关接口
package service

import (
	"backEnd/manager"
	"backEnd/types"
	"log"

	"github.com/gin-gonic/gin"
)

// PostAmountByYearAndMonthFromOnlineOrderStatisticService
//
//	@Description: 获取网上订单量接口
//	@param c
func PostAmountByYearAndMonthFromOnlineOrderStatisticService(c *gin.Context) {
	var request types.PostAmountByYearAndMonthFromOnlineOrderStatisticRequest
	if err := c.ShouldBind(&request); err != nil {
		log.Println(err)
		return
	}

	if request.BeginYear < types.OrderMinYear {
		request.BeginYear = types.OrderMinYear
	}

	if request.EndYear < types.OrderMaxYear {
		request.EndYear = types.OrderMaxYear
	}

	statisticManager := manager.OnlineOrderStatisticManager{}
	allTime, allAmount, errNo := statisticManager.PostAmountByYearAndMonthFromOnlineOrderStatistic(request.BeginYear, request.EndYear)
	if errNo != types.OK {
		c.JSON(200, types.PostAmountByYearAndMonthFromOnlineOrderStatisticResponse{Code: errNo})
	} else {
		c.JSON(200, types.PostAmountByYearAndMonthFromOnlineOrderStatisticResponse{
			Code: types.OK,
			Data: struct {
				AllTime   []string `json:"allTime"`
				AllAmount []string `json:"allAmount"`
			}{AllTime: allTime, AllAmount: allAmount},
		})
	}
}
