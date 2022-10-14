// Package dao
// @Description: 市场规模 dao 类
package dao

import (
	"backEnd/types"
	"fmt"
	"log"
	"strconv"
)

type StatisticDao struct {
}

// SelectAmountAndPercentageByYearFromMySQL
//
//	@Description: 从 MySQL 中通过 year 获取市场规模和百分比后写入 redis
//	@receiver s *StatisticDao
//	@param year
//	@return types.ErrNo
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

// InsertAmountAndPercentageToRedis
//
//	@Description: 将市场规模和百分比按年份写入 Redis
//	@receiver s *StatisticDao
//	@param year
//	@param amount
//	@param percentage
//	@return types.ErrNo
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

// SelectAmountAndPercentageByYearFromRedis
//
//	@Description: 从 Redis 中按年份获取市场规模和百分比
//	@receiver s *StatisticDao
//	@param year
//	@return string
//	@return string
//	@return types.ErrNo
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
