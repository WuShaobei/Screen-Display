// Package manager
// @Description: 融资数据 manager 类
package manager

import (
	"backEnd/dao"
	"backEnd/types"
	"sort"
)

type FundingStatisticManager struct {
	fundingStatisticDao dao.FundingStatisticDao
}

// GetAllDataFromFundingStatistic
//
//	@Description: 获取所有融资数据
//	@receiver f *FundingStatisticManager
//	@return []types.GetAllDataFromFundingStatisticData
//	@return types.ErrNo
func (f *FundingStatisticManager) GetAllDataFromFundingStatistic() ([]types.GetAllDataFromFundingStatisticData, types.ErrNo) {
	data, errNo := f.fundingStatisticDao.SelectAllDataFromMysql()
	if errNo != types.OK {
		return data, errNo
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i].Time > data[j].Time
	})
	return data, errNo
}
