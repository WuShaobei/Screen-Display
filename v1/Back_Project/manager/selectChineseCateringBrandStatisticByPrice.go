package manager

import (
	"Back_Project/dao"
	"strings"
)

// SelectChineseCateringBrandStatisticByPrice
//
//	@Description: 获取价格区间段中的 brand 数据 [minPrice, maxPrice)
//	@param minPrice
//	@param maxPrice
//	@return string
//	@data 2022-09-28 13:43:28
//	@author WuShaobei
func SelectChineseCateringBrandStatisticByPrice(minPrice, maxPrice int) string {

	var sqlData []string

	sqlStr := `
		SELECT brand
		From chinese_catering_brand_statistics
		WHERE ? <= price AND price < ? 
	`
	dao.DB.Raw(sqlStr, minPrice, maxPrice).Scan(&sqlData)

	return strings.Join(sqlData, "\n")
}
