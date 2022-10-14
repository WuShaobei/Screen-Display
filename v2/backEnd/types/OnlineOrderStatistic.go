// Package types
// @Description: 网上订单量相关类

package types

// ChineseCateringOnlineOrderStatistic
// @Description: 订单量 Gorm模型
type ChineseCateringOnlineOrderStatistic struct {
	Id          uint `gorm:"primaryKey BIGINT NOT NULL AUTOINCREMENT"`
	Year        int  `gorm:"INT NULL DEFAULT NULL"`
	Month       int  `gorm:"INT NULL DEFAULT NULL"`
	OrderAmount int  `gorm:"INT NULL DEFAULT NULL"`
}

// PostAmountByYearAndMonthFromOnlineOrderStatisticRequest
// @Description: 请求订单量接口请求类型
type PostAmountByYearAndMonthFromOnlineOrderStatisticRequest struct {
	BeginYear int `json:"beginYear"`
	EndYear   int `json:"endYear"`
}

// PostAmountByYearAndMonthFromOnlineOrderStatisticResponse
// @Description: 请求订单量接口返回内容
type PostAmountByYearAndMonthFromOnlineOrderStatisticResponse struct {
	Code ErrNo `json:"code"`
	Data struct {
		AllTime   []string `json:"allTime"`
		AllAmount []string `json:"allAmount"`
	} `json:"data"`
}
