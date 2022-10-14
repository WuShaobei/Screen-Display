package dao

import (
	"backEnd/types"
	"log"
	"strconv"
	"strings"
)

type BrandStatisticDao struct {
}

func (b *BrandStatisticDao) SelectDataByPriceFromMySQL(minPrice, maxPrice int) types.ErrNo {
	var data []string
	DB.Table("chinese_catering_brand_statistics").Where(" ? <= price AND Price < ?", minPrice, maxPrice).Select("brand").Find(&data)
	return b.InsertDataByPriceToRedis(
		strconv.Itoa(minPrice),
		strconv.Itoa(maxPrice),
		strings.Join(data, "\n"),
	)
}

func (b *BrandStatisticDao) InsertDataByPriceToRedis(minPrice, maxPrice, data string) types.ErrNo {
	key := "BrandStatisticDataFrom" + minPrice + "To" + maxPrice
	if err := RDB.HSet(key, "data", data).Err(); err != nil {
		log.Fatal(err)
		return types.UnknownError
	}
	return types.OK
}

func (b *BrandStatisticDao) SelectDataByPriceFromRedis(minPrice, maxPrice string) (string, types.ErrNo) {
	key := "BrandStatisticDataFrom" + minPrice + "To" + maxPrice
	data, err := RDB.HGet(key, "data").Result()
	if err != nil {
		return "", types.UserNotExist
	}
	return data, types.OK
}
