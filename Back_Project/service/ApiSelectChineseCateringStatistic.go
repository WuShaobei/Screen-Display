package service

import (
	"Back_Project/manager"
	"Back_Project/types"
	"log"

	"github.com/gin-gonic/gin"
)

// SelectChineseCateringStatistic
//
//	@Description: 获取 request.BeginYear - request.EndYear 中每年的市场份额和百分比
//	@param c
//	@data 2022-09-28 14:10:08
//	@author WuShaobei
func SelectChineseCateringStatistic(c *gin.Context) {
	var request types.SelectChineseCateringStatisticRequest
	if err := c.ShouldBind(&request); err != nil {
		log.Println(err)
		return
	}

	var response types.SelectChineseCateringStatisticResponse
	for year := request.BeginYear; year <= request.EndYear; year++ {
		response.Data = append(response.Data, manager.SelectChineseCateringStatisticByYear(year))
	}

	c.JSON(200, response)

}
