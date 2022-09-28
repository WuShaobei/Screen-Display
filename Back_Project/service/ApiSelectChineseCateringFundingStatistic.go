package service

import (
	"Back_Project/manager"
	"Back_Project/types"

	"github.com/gin-gonic/gin"
)

// SelectChineseCateringFundingStatistic
//
//	@Description: 逆序获取投资数据并返回
//	@param c
//	@data 2022-09-28 14:04:55
//	@author WuShaobei
func SelectChineseCateringFundingStatistic(c *gin.Context) {

	var response types.SelectChineseCateringFundingStatisticResponse
	for year := 2022; year >= 2021; year-- {
		for month := 12; month >= 1; month-- {
			res := manager.SelectChineseCateringFundingStatisticByYearAndMonth(year, month)
			response.Data = append(response.Data, res.Data...)
		}
	}

	c.JSON(200, response)

}
