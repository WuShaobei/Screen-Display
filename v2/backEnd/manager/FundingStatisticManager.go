package manager

import (
	"backEnd/dao"
	"backEnd/types"
	"sort"
)

type FundingStatisticManager struct {
	fundingStatisticDao dao.FundingStatisticDao
}

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
