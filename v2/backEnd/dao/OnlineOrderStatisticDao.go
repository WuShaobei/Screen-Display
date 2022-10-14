// Package dao
// @Description: 点单数据 dao 类

package dao

import (
	"backEnd/types"
	"log"
	"strconv"
)

type OnlineOrderStatisticDao struct {
}

// SelectAmountFromYearAndMonthFromMySQL
//
//	@Description: 从 MySQL 中获取 year 年 month 月的订单数然后写入 redis
//	@receiver o *OnlineOrderStatisticDao
//	@param year
//	@param month
//	@return types.ErrNo
//	@data 2022-10-14 16:15:00
//	@author WuShaobei
func (o *OnlineOrderStatisticDao) SelectAmountFromYearAndMonthFromMySQL(year, month int) types.ErrNo {
	var amount string
	DB.Table("chinese_catering_online_order_statistics").Where("year=? AND month = ?", year, month).Select("order_amount").Find(&amount)

	return o.InsertAmountFromYearAndMonthToRedis(
		strconv.Itoa(year)+"-"+strconv.Itoa(month),
		amount,
	)
}
func (o *OnlineOrderStatisticDao) InsertAmountFromYearAndMonthToRedis(date, amount string) types.ErrNo {
	key := "OnlineOrderStatisticDataOn" + date
	if err := RDB.HSet(key, "amount", amount).Err(); err != nil {
		log.Fatal(err)
		return types.UnknownError
	}
	return types.OK
}

func (o *OnlineOrderStatisticDao) SelectAmountFromYearAndMonthFromRedis(date string) (string, types.ErrNo) {
	key := "OnlineOrderStatisticDataOn" + date
	if amount, err := RDB.HGet(key, "amount").Result(); err == nil {
		return amount, types.OK
	} else {
		return "", types.UnknownError
	}
}
