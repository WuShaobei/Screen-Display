// Package service
// @Description: 火锅品牌相关接口
package service

import (
	"backEnd/manager"
	"backEnd/types"

	"github.com/gin-gonic/gin"
)

// GetAllDataFromBrandStatisticService
//
//	@Description: 获取所有火锅品牌信息
//	@param c
func GetAllDataFromBrandStatisticService(c *gin.Context) {
	statisticManager := manager.BrandStatisticManager{}
	if data, errNo := statisticManager.GetAllData(); errNo != types.OK {
		c.JSON(200,
			types.GetAllDataFromBrandStatisticResponse{
				Code: errNo,
			})
	} else {
		c.JSON(200,
			types.GetAllDataFromBrandStatisticResponse{
				Code: errNo,
				Data: data,
			})
	}

}
