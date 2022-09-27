package service

import (
	"Back_Project/manager"
	"Back_Project/types"
	"log"

	"github.com/gin-gonic/gin"
)

func Api2(c *gin.Context) {
	var request types.Api2Request
	if err := c.ShouldBind(&request); err != nil {
		log.Println(err)
		return
	}

	var response types.Api2Response
	for year := request.BeginYear; year <= request.EndYear; year++ {
		for month := 1; month <= 12; month++ {
			res := manager.GetChineseCateringOnlineOrderStatistic(year, month)
			response.AllTime = append(response.AllTime, res.Time)
			response.AllOrderAmount = append(response.AllOrderAmount, res.OrderAmount)
		}
	}

	c.JSON(200, response)

}
