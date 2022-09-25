package manager

import (
	"Back_Project/dao"
	"Back_Project/types"
)

func GetChineseCateringPayment(year int) types.Api3Data {
	var sqlData types.Api3Data
	sqlData.Year = year

	sqlStr := `
		SELECT AVG( salary ) AS avg_salary, COUNT(*) AS count_salary
		From chinese_catering_payments
		WHERE year = ?
	`
	dao.DB.Raw(sqlStr, year).Scan(&sqlData)

	return sqlData
}
