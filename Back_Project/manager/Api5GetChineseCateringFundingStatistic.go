package manager

import (
	"Back_Project/dao"
	"Back_Project/types"
	"strconv"
)

func GetChineseCateringFundingStatistic(year, month int) types.Api5Response {
	var sqlData []types.ChineseCateringFundingStatistic

	sqlStr := `
		SELECT *
		From chinese_catering_funding_statistics
		WHERE funding_year = ? AND funding_month = ?
	`
	dao.DB.Raw(sqlStr, year, month).Scan(&sqlData)
	str := strconv.Itoa(year) + "-" + strconv.Itoa(month)
	var res types.Api5Response
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
