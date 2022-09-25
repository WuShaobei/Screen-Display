package manager

import (
	"Back_Project/dao"
	"Back_Project/types"
)

func GetChineseCateringFundingStatistic(year, month int) types.Api5Response {
	var sqlData types.Api5Response

	sqlStr := `
		SELECT brand AS Brand, funding_round AS FundingRound, investor AS Investor
		From chinese_catering_funding_statistics
		WHERE funding_year = ? AND funding_month = ?
	`
	dao.DB.Raw(sqlStr, year, month).Scan(&sqlData)

	return sqlData
}
