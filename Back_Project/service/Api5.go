package service

import (
	"Back_Project/manager"
	"Back_Project/types"

	"github.com/gin-gonic/gin"
)

func Api5(c *gin.Context) {

	var response types.Api5Response
	for year := 2022; year >= 2021; year-- {
		for month := 12; month >= 1; month-- {

			res := manager.GetChineseCateringFundingStatistic(year, month)
			for _, data := range res.Data {
				response.Data = append(response.Data, data)
			}
		}
	}

	c.JSON(200, response)

}
