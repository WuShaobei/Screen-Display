package service

import (
	"backEnd/manager"
	"backEnd/types"

	"github.com/gin-gonic/gin"
)

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
