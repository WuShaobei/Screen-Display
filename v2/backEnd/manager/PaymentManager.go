// Package manager
// @Description: 薪资 manager 类
package manager

import (
	"backEnd/dao"
	"backEnd/types"
	"strconv"
)

type PaymentManager struct {
	paymentDao *dao.PaymentDao
}

// PostAvgSalaryAndCountsByYearFromPaymentManage
//
//	@Description: 获取平均薪资和人数
//	@receiver p *PaymentManager
//	@param beginYear
//	@param endYear
//	@return []types.PostAvgSalaryAndCountsByYearFromPaymentData
//	@return types.ErrNo
func (p *PaymentManager) PostAvgSalaryAndCountsByYearFromPaymentManage(beginYear, endYear int) ([]types.PostAvgSalaryAndCountsByYearFromPaymentData, types.ErrNo) {
	need := 0
	have := 0
	var res []types.PostAvgSalaryAndCountsByYearFromPaymentData

	for year := beginYear; year <= endYear; year++ {
		if data, errNo := p.paymentDao.SelectAvgSalaryAndCountsByYearFromRedis(strconv.Itoa(year)); errNo == types.OK {
			res = append(res, data)
			have += 1
			need += 1
		} else {
			go p.paymentDao.SelectAvgSalaryAndCountsByYearFromMySQL(year)
			res = append(res, types.PostAvgSalaryAndCountsByYearFromPaymentData{})
			need += 1
		}
	}

	for have < need {
		idx := 0
		for year := beginYear; year <= endYear; year++ {
			if res[idx] == (types.PostAvgSalaryAndCountsByYearFromPaymentData{}) {
				if data, errNo := p.paymentDao.SelectAvgSalaryAndCountsByYearFromRedis(strconv.Itoa(year)); errNo == types.OK {
					res[idx] = data
					have += 1
				}
			} else {
				continue
			}
		}
	}

	return res, types.OK
}
