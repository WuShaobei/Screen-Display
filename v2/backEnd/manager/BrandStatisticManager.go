package manager

import (
	"backEnd/dao"
	"backEnd/types"
	"strconv"
)

type BrandStatisticManager struct {
	brandStatisticDao dao.BrandStatisticDao
}

func (b *BrandStatisticManager) GetData() ([]string, types.ErrNo) {
	var res []string
	need := 0
	have := 0
	var prices = [...][2]int{{0, 80}, {80, 120}, {120, 200}, {200, 10000000}}

	for _, price := range prices {
		if data, errNo := b.brandStatisticDao.SelectDataByPriceFromRedis(
			strconv.Itoa(price[0]),
			strconv.Itoa(price[1]),
		); errNo != types.OK {
			go b.brandStatisticDao.SelectDataByPriceFromMySQL(price[0], price[1])
			res = append(res, "")
			need += 1
		} else {
			res = append(res, data)
			need += 1
			have += 1
		}
	}

	for have < need {
		for i, price := range prices {
			if data, errNo := b.brandStatisticDao.SelectDataByPriceFromRedis(
				strconv.Itoa(price[0]),
				strconv.Itoa(price[1]),
			); errNo != types.OK {
				continue
			} else {
				res[i] = data
				have += 1
			}
		}
	}

	return res, types.OK
}
