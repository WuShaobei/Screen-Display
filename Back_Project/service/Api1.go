package service

import (
	"Back_Project/manager"
	"Back_Project/types"
	"log"

	"github.com/gin-gonic/gin"
)

func Api1(c *gin.Context) {
	var request types.Api1Request
	if err := c.ShouldBind(&request); err != nil {
		log.Println(err)
		return
	}

	var response types.Api1Response
	for year := request.BeginYear; year <= request.EndYear; year++ {
		response.Data = append(response.Data, manager.GetChineseCateringStatistic(year))
	}

	c.JSON(200, response)

}
