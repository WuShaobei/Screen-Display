package manager

import (
	"backEnd/dao"
	"backEnd/types"
	"strconv"
)

type OnlineOrderStatisticManager struct {
	onlineOrderStatisticDao dao.OnlineOrderStatisticDao
}

func (o *OnlineOrderStatisticManager) PostAmountByYearAndMonthFromOnlineOrderStatistic(beginYear, endYear int) ([]string, []string, types.ErrNo) {
	var allTime []string
	var allAmount []string

	have := 0
	need := 0

	for year := beginYear; year <= endYear; year++ {
		for month := 1; month <= 12; month++ {
			date := strconv.Itoa(year) + "-" + strconv.Itoa(month)
			allTime = append(allTime, date)
			need += 1
			if amount, errNo := o.onlineOrderStatisticDao.SelectAmountFromYearAndMonthFromRedis(date); errNo == types.OK {
				allAmount = append(allAmount, amount)
				have += 1
			} else {
				allAmount = append(allAmount, "")
				go o.onlineOrderStatisticDao.SelectAmountFromYearAndMonthFromMySQL(year, month)
			}
		}
	}

	for have < need {
		idx := 0
		for year := beginYear; year <= endYear; year++ {
			for month := 1; month <= 12; month++ {
				if allAmount[idx] != "" {
					idx += 1
					continue
				}
				date := strconv.Itoa(year) + "-" + strconv.Itoa(month)
				if amount, errNo := o.onlineOrderStatisticDao.SelectAmountFromYearAndMonthFromRedis(date); errNo == types.OK {

					allAmount[idx] = amount
					have += 1
				}
			}
		}
	}

	return allTime, allAmount, types.OK
}
