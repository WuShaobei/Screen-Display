// Package types
// @Description: 火锅品牌相关类

package types

// ChineseCateringBrandStatistic
// @Description: 火锅品牌Gorm模型
type ChineseCateringBrandStatistic struct {
	Id    uint   `gorm:"primaryKey BIGINT NOT NULL AUTOINCREMENT"`
	Brand string `gorm:"VARCHAR(255) DEFAULT NULL"`
	Price int    `gorm:"INT NULL DEFAULT NULL"`
	Other string `gorm:"VARCHAR(255) NULL DEFAULT NULL"`
}

// GetAllDataFromBrandStatisticResponse
// @Description: 获取所有火锅品牌数据接口返回内容
type GetAllDataFromBrandStatisticResponse struct {
	Code ErrNo    `json:"code"`
	Data []string `json:"data"`
}
