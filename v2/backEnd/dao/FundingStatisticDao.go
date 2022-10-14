// Package dao
// @Description: 融资数据 dao 类
package dao

import "backEnd/types"

type FundingStatisticDao struct {
}

// SelectAllDataFromMysql
//
//	@Description: 从 MySQL 中获取所有融资数据
//	@receiver f *FundingStatisticDao
//	@return []types.GetAllDataFromFundingStatisticData
//	@return types.ErrNo
func (f *FundingStatisticDao) SelectAllDataFromMysql() ([]types.GetAllDataFromFundingStatisticData, types.ErrNo) {

	var allData []types.ChineseCateringFundingStatistic
	DB.Find(&allData)

	var res []types.GetAllDataFromFundingStatisticData
	for _, data := range allData {
		res = append(res, types.GetAllDataFromFundingStatisticData{
			Time:         types.GetTime(data.FundingYear, data.FundingMonth),
			Brand:        data.Brand,
			FundingRound: data.FundingRound,
			Investor:     data.Investor,
		})
	}

	return res, types.OK
}
