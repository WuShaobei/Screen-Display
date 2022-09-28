package manager

import (
	"Back_Project/dao"
	"Back_Project/types"
	"strconv"
)

// SelectChineseCateringFundingStatisticByYearAndMonth
//
//	@Description: 获取 year 年 month 月所有的 品牌、融资轮次、投资商
//	@param year
//	@param month
//	@return types.SelectChineseCateringFundingStatisticResponse
//	@data 2022-09-28 13:44:54
//	@author WuShaobei
func SelectChineseCateringFundingStatisticByYearAndMonth(year, month int) types.SelectChineseCateringFundingStatisticResponse {
	var sqlData []types.ChineseCateringFundingStatistic

	sqlStr := `
		SELECT *
		From chinese_catering_funding_statistics
		WHERE funding_year = ? AND funding_month = ?
	`
	dao.DB.Raw(sqlStr, year, month).Scan(&sqlData)
	str := strconv.Itoa(year) + "-" + strconv.Itoa(month)
	var res types.SelectChineseCateringFundingStatisticResponse
	for _, data := range sqlData {
		res.Data = append(res.Data, struct {
			Time         string
			Brand        string
			FundingRound string
			Investor     string
		}{Time: str, Brand: data.Brand, FundingRound: data.FundingRound, Investor: data.Investor})
	}

	return res
}
