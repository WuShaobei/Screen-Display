package manager

import (
	"Back_Project/dao"
	"strings"
)

func GetChineseCateringBrandStatistic(minNum, maxNum int) string {

	var sqlData []string

	sqlStr := `
		SELECT brand
		From chinese_catering_brand_statistics
		WHERE ? <= price AND price < ? 
	`
	dao.DB.Raw(sqlStr, minNum, maxNum).Scan(&sqlData)

	return strings.Join(sqlData, "\n")
}
