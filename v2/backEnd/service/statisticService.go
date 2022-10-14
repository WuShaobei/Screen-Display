package service

import (
	"backEnd/manager"
	"backEnd/types"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func PostAmountAndPercentageByYearFromStatisticService(c *gin.Context) {
	var request types.PostAmountAndPercentageByYearFromStatisticRequest
	if err := c.ShouldBind(&request); err != nil {
		log.Println(err)
		return
	}

	if request.BeginYear < types.StatisticMinYear {
		request.BeginYear = types.StatisticMinYear
	}

	if request.EndYear < types.StatisticMaxYear {
		request.EndYear = types.StatisticMaxYear
	}
	statisticManager := manager.StatisticManager{}

	data, errNo := statisticManager.PostAmountAndPercentageByYear(request.BeginYear, request.EndYear)
	fmt.Println(data)
	if errNo == types.OK {
		c.JSON(200, types.PostAmountAndPercentageByYearFromStatisticResponse{
			Code: errNo,
			Data: data,
		})
	} else {
		c.JSON(200, types.PostAmountAndPercentageByYearFromStatisticResponse{
			Code: errNo,
		})
	}

}
