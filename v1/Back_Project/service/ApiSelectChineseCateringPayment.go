package service

import (
	"Back_Project/manager"
	"Back_Project/types"
	"log"

	"github.com/gin-gonic/gin"
)

// SelectChineseCateringPayment
//
//	@Description: 获取 request.BeginYear - request.EndYear 中每年的人数和平均报酬
//	@param c
//	@data 2022-09-28 14:07:56
//	@author WuShaobei
func SelectChineseCateringPayment(c *gin.Context) {
	var request types.SelectGetChineseCateringPaymentRequest
	if err := c.ShouldBind(&request); err != nil {
		log.Println(err)
		return
	}

	var response types.SelectGetChineseCateringPaymentResponse
	for year := request.BeginYear; year <= request.EndYear; year++ {
		response.Data = append(response.Data, manager.SelectChineseCateringPaymentByYear(year))
	}

	c.JSON(200, response)

}
