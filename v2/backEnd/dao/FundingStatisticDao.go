package dao

import "backEnd/types"

type FundingStatisticDao struct {
}

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
