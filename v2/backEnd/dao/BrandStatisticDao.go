// Package dao
// @Description: 火锅品牌 dao 类
package dao

import (
	"backEnd/types"
	"log"
	"strconv"
	"strings"
)

type BrandStatisticDao struct {
}

// SelectDataByPriceFromMySQL
//
//	@Description: 从 MySQL 中获取对应价格段的火锅品牌数据然后写入 redis
//	@receiver b *BrandStatisticDao
//	@param minPrice
//	@param maxPrice
//	@return types.ErrNo
func (b *BrandStatisticDao) SelectDataByPriceFromMySQL(minPrice, maxPrice int) types.ErrNo {
	var data []string
	DB.Table("chinese_catering_brand_statistics").Where(" ? <= price AND Price < ?", minPrice, maxPrice).Select("brand").Find(&data)
	return b.InsertDataByPriceToRedis(
		strconv.Itoa(minPrice),
		strconv.Itoa(maxPrice),
		strings.Join(data, "\n"),
	)
}

// InsertDataByPriceToRedis
//
//	@Description: 将价格段的品牌插入 redis
//	@receiver b *BrandStatisticDao
//	@param minPrice
//	@param maxPrice
//	@param data
//	@return types.ErrNo
func (b *BrandStatisticDao) InsertDataByPriceToRedis(minPrice, maxPrice, data string) types.ErrNo {
	key := "BrandStatisticDataFrom" + minPrice + "To" + maxPrice
	if err := RDB.HSet(key, "data", data).Err(); err != nil {
		log.Fatal(err)
		return types.UnknownError
	}
	return types.OK
}

// SelectDataByPriceFromRedis
//
//	@Description:  从 redis 中获取对应价格段的火锅品牌
//	@receiver b
//	@param minPrice
//	@param maxPrice
//	@return string
//	@return types.ErrNo
func (b *BrandStatisticDao) SelectDataByPriceFromRedis(minPrice, maxPrice string) (string, types.ErrNo) {
	key := "BrandStatisticDataFrom" + minPrice + "To" + maxPrice
	data, err := RDB.HGet(key, "data").Result()
	if err != nil {
		return "", types.UserNotExist
	}
	return data, types.OK
}
