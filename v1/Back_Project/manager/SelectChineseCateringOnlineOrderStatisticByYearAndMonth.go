package manager

import (
	"Back_Project/dao"
	"Back_Project/types"
	"strconv"
)

func SelectChineseCateringOnlineOrderStatisticByYearAndMonth(year, month int) types.SelectChineseCateringOnlineOrderStatisticData {
	var sqlData types.SelectChineseCateringOnlineOrderStatisticData
	sqlData.Time = strconv.Itoa(year) + "-" + strconv.Itoa(month)

	sqlStr := `
		SELECT order_amount AS OrderAmount
		From chinese_catering_online_order_statistics
		WHERE year = ? AND month = ?
	`
	dao.DB.Raw(sqlStr, year, month).Scan(&sqlData)

	return sqlData
}
