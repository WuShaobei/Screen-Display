package dao

import (
	"backEnd/types"
	"log"
	"strconv"
)

type OnlineOrderStatisticDao struct {
}

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
