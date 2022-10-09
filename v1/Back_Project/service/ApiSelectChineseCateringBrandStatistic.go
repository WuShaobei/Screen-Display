package service

import (
	"Back_Project/manager"
	"Back_Project/types"

	"github.com/gin-gonic/gin"
)

// SelectChineseCateringBrandStatistic
//
//	@Description: 返回 四个档次 的所有 品牌
//	@param c
//	@data 2022-09-28 14:01:57
//	@author WuShaobei
func SelectChineseCateringBrandStatistic(c *gin.Context) {
	response := struct {
		Code types.ErrNo
		Data []string
	}{}

	response.Data = append(response.Data, manager.SelectChineseCateringBrandStatisticByPrice(0, 80))
	response.Data = append(response.Data, manager.SelectChineseCateringBrandStatisticByPrice(80, 120))
	response.Data = append(response.Data, manager.SelectChineseCateringBrandStatisticByPrice(120, 200))
	response.Data = append(response.Data, manager.SelectChineseCateringBrandStatisticByPrice(200, 1000))

	c.JSON(200, response)
}
