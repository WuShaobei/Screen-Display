package service

import (
	"Back_Project/manager"
	"Back_Project/types"
	"log"

	"github.com/gin-gonic/gin"
)

// SelectChineseCateringOnlineOrderStatistic
//
//	@Description: 获取 request.beginYear - request.EndYear 中每个月的订单量
//	@param c
//	@data 2022-09-28 14:05:29
//	@author WuShaobei
func SelectChineseCateringOnlineOrderStatistic(c *gin.Context) {
	var request types.SelectChineseCateringOnlineOrderStatisticRequest
	if err := c.ShouldBind(&request); err != nil {
		log.Println(err)
		return
	}

	var response types.SelectChineseCateringOnlineOrderStatisticResponse
	for year := request.BeginYear; year <= request.EndYear; year++ {
		for month := 1; month <= 12; month++ {
			res := manager.SelectChineseCateringOnlineOrderStatisticByYearAndMonth(year, month)
			response.AllTime = append(response.AllTime, res.Time)
			response.AllOrderAmount = append(response.AllOrderAmount, res.OrderAmount)
		}
	}

	c.JSON(200, response)

}
