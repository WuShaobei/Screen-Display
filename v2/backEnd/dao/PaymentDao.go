package dao

import (
	"backEnd/types"
	"log"
	"strconv"
)

type PaymentDao struct {
}

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
