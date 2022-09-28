package manager

import (
	"Back_Project/dao"
	"Back_Project/types"
)

// SelectChineseCateringPaymentByYear
//
//	@Description: 获取 year 年的 平均薪资 和 人数
//	@param year
//	@return types.SelectGetChineseCateringPaymentData
//	@data 2022-09-28 13:34:10
func SelectChineseCateringPaymentByYear(year int) types.SelectGetChineseCateringPaymentData {
	var sqlData types.SelectGetChineseCateringPaymentData
	sqlData.Year = year

	sqlStr := `
		SELECT AVG( salary ) AS avg_salary, COUNT(*) AS count_salary
		From chinese_catering_payments
		WHERE year = ?
	`
	dao.DB.Raw(sqlStr, year).Scan(&sqlData)
	sqlData.AvgSalary = dao.Decimal(sqlData.AvgSalary)
	return sqlData
}
