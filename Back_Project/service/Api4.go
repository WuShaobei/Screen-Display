package service

import (
	"Back_Project/manager"
	"Back_Project/types"

	"github.com/gin-gonic/gin"
)

func Api4(c *gin.Context) {
	response := struct {
		Code types.ErrNo
		Data []string
	}{}

	response.Data = append(response.Data, manager.GetChineseCateringBrandStatistic(0, 80))
	response.Data = append(response.Data, manager.GetChineseCateringBrandStatistic(80, 120))
	response.Data = append(response.Data, manager.GetChineseCateringBrandStatistic(120, 200))
	response.Data = append(response.Data, manager.GetChineseCateringBrandStatistic(200, 1000))

	c.JSON(200, response)
}
