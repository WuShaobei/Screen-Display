// Package service
// @Description: 融资相关接口
package service

import (
	"backEnd/manager"
	"backEnd/types"

	"github.com/gin-gonic/gin"
)

// GetAllDataByFromFundingStatisticService
//
//	@Description: 获取所有融资信息接口
//	@param c
func GetAllDataByFromFundingStatisticService(c *gin.Context) {

	fundingStatisticManager := manager.FundingStatisticManager{}
	if data, errNo := fundingStatisticManager.GetAllDataFromFundingStatistic(); errNo == types.OK {
		c.JSON(200, types.PostAllDataByYearAndMonthFromFundingStatisticResponse{
			Code: types.OK,
			Data: data,
		})
	} else {
		c.JSON(200, types.PostAllDataByYearAndMonthFromFundingStatisticResponse{Code: errNo})
	}
}
