package service

import (
	"backEnd/manager"
	"backEnd/types"

	"github.com/gin-gonic/gin"
)

func GetAllDataFromBrandStatisticService(c *gin.Context) {
	statisticManager := manager.BrandStatisticManager{}
	if data, errNo := statisticManager.GetData(); errNo != types.OK {
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
