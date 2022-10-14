package dao

import (
	"backEnd/types"
	"fmt"
	"log"
	"strconv"
)

type StatisticDao struct {
	//data types.ChineseCateringStatistic
}

func (s *StatisticDao) SelectAmountAndPercentageByYearFromMySQL(year int) types.ErrNo {
	var data types.ChineseCateringStatistic
	DB.Where("year = ?", year).Find(&data)
	fmt.Println(data)
	errNo := s.InsertAmountAndPercentageToRedis(
		strconv.Itoa(year), types.Decimal(data.TotalAmount), types.Decimal(data.TotalAmountPercentage),
	)
	if errNo != types.OK {
		return errNo
	}
	return errNo
}

func (s *StatisticDao) InsertAmountAndPercentageToRedis(year, amount, percentage string) types.ErrNo {
	if err := RDB.HSet("StatisticDataIn"+year, "amount", amount).Err(); err != nil {
		log.Fatal(err)
		return types.UnknownError
	}
	if err := RDB.HSet("StatisticDataIn"+year, "percentage", percentage).Err(); err != nil {
		log.Fatal(err)
		return types.UnknownError
	}
	return types.OK
}

func (s *StatisticDao) SelectAmountAndPercentageByYearFromRedis(year string) (string, string, types.ErrNo) {
	amount, err := RDB.HGet("StatisticDataIn"+year, "amount").Result()
	if err != nil {
		return "", "", types.UserNotExist
	}
	percentage, err := RDB.HGet("StatisticDataIn"+year, "percentage").Result()
	if err != nil {
		return "", "", types.UserNotExist
	}
	return amount, percentage, types.OK
}
