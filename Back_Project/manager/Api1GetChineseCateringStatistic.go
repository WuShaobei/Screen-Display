package manager

import (
	"Back_Project/dao"
	"Back_Project/types"
)

func GetChineseCateringStatistic(year int) types.Api1Data {
	var sqlData types.Api1Data
	sqlData.Year = year

	sqlStr := `
		SELECT total_amount, total_amount_percentage  
		From chinese_catering_statistics
		WHERE year = ?
	`
	dao.DB.Raw(sqlStr, year).Scan(&sqlData)

	return sqlData
}
