// Package manager
// @Description: 餐饮市场 manager 类

package manager

import (
	"backEnd/dao"
	"backEnd/types"
	"strconv"
)

type StatisticManager struct {
	statisticDao dao.StatisticDao
}

// PostAmountAndPercentageByYear
//
//	@Description: 获取规模和占比
//	@receiver s StatisticManager
//	@param beginYear
//	@param endYear
//	@return []types.PostAmountAndPercentageByYearFromStatisticData
//	@return types.ErrNo
func (s *StatisticManager) PostAmountAndPercentageByYear(beginYear, endYear int) ([]types.PostAmountAndPercentageByYearFromStatisticData, types.ErrNo) {
	need := endYear - beginYear + 1
	res := make([]types.PostAmountAndPercentageByYearFromStatisticData, need)
	have := 0

	for i := 0; i < need; i++ {
		if (res[i] != types.PostAmountAndPercentageByYearFromStatisticData{}) {
			continue
		}
		amount, percentage, errNo := s.statisticDao.SelectAmountAndPercentageByYearFromRedis(strconv.Itoa(i + beginYear))
		if errNo == types.OK {
			res[i].Year = strconv.Itoa(i + beginYear)
			res[i].Amount = amount
			res[i].Percentage = percentage
			have += 1
		} else {
			go s.statisticDao.SelectAmountAndPercentageByYearFromMySQL(i + beginYear)
		}
	}

	for have < need {
		for i := 0; i < need; i++ {
			if (res[i] != types.PostAmountAndPercentageByYearFromStatisticData{}) {
				continue
			}
			amount, percentage, errNo := s.statisticDao.SelectAmountAndPercentageByYearFromRedis(strconv.Itoa(i + beginYear))
			if errNo == types.OK {
				res[i].Year = strconv.Itoa(i + beginYear)
				res[i].Amount = amount
				res[i].Percentage = percentage
				have += 1
			} else {
				continue
			}
		}
	}

	return res, types.OK
}
