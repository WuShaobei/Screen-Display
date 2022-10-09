package manager

import (
	"Back_Project/dao"
	"Back_Project/types"
)

// SelectChineseCateringStatisticByYear
//
//	@Description: 获取 year 年的 总金额 和 所占比例
//	@param year
//	@return types.SelectChineseCateringStatisticData
//	@data 2022-09-28 13:49:09
//	@author WuShaobei
func SelectChineseCateringStatisticByYear(year int) types.SelectChineseCateringStatisticData {
	var sqlData types.SelectChineseCateringStatisticData
	sqlData.Year = year

	sqlStr := `
		SELECT total_amount, total_amount_percentage  
		From chinese_catering_statistics
		WHERE year = ?
	`
	dao.DB.Raw(sqlStr, year).Scan(&sqlData)

	return sqlData
}
