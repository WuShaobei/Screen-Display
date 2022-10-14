// Package dao
// @Description: 薪资 dao 类

package dao

import (
	"backEnd/types"
	"log"
	"strconv"
)

type PaymentDao struct {
}

// SelectAvgSalaryAndCountsByYearFromMySQL
//
//	@Description: 从 MySQL 中获取 year 年的平均薪资和人数然后写入 redis
//	@receiver p *PaymentDao
//	@param year
//	@return types.ErrNo
func (p *PaymentDao) SelectAvgSalaryAndCountsByYearFromMySQL(year int) types.ErrNo {
	data := struct {
		Avg    float64
		Counts string
	}{}
	DB.Table("chinese_catering_payments").Where("year = ?", year).Select("AVG(salary) as avg, COUNT(salary) as counts").Find(&data)
	if data == (struct {
		Avg    float64
		Counts string
	}{}) {
		return types.UnknownError
	}
	return p.insertAvgSalaryAndCountsDataToRedis(strconv.Itoa(year), types.Decimal(data.Avg), data.Counts)
}

// insertAvgSalaryAndCountsDataToRedis
//
//	@Description: 将 year 年的平均薪资和人数写入 redis
//	@receiver p *PaymentDao
//	@param year
//	@param avgSalary
//	@param counts
//	@return types.ErrNo
func (p *PaymentDao) insertAvgSalaryAndCountsDataToRedis(year, avgSalary, counts string) types.ErrNo {
	if err := RDB.HSet("PaymentDataIn"+year, "avgSalary", avgSalary).Err(); err != nil {
		log.Fatal(err)
		return types.UnknownError
	}
	if err := RDB.HSet("PaymentDataIn"+year, "counts", counts).Err(); err != nil {
		log.Fatal(err)
		return types.UnknownError
	}
	return types.OK
}

// SelectAvgSalaryAndCountsByYearFromRedis
//
//	@Description: 从 redis 中获取 year 年的平均薪资和人数
//	@receiver p *PaymentDao
//	@param year
//	@return types.PostAvgSalaryAndCountsByYearFromPaymentData
//	@return types.ErrNo
func (p *PaymentDao) SelectAvgSalaryAndCountsByYearFromRedis(year string) (types.PostAvgSalaryAndCountsByYearFromPaymentData, types.ErrNo) {
	data := types.PostAvgSalaryAndCountsByYearFromPaymentData{Year: year}

	var err error
	data.AvgSalary, err = RDB.HGet("PaymentDataIn"+year, "avgSalary").Result()
	if err != nil {
		return types.PostAvgSalaryAndCountsByYearFromPaymentData{}, types.UnknownError
	}
	data.Counts, err = RDB.HGet("PaymentDataIn"+year, "counts").Result()
	if err != nil {
		return types.PostAvgSalaryAndCountsByYearFromPaymentData{}, types.UnknownError
	}
	return data, types.OK
}
