// Package types
// @Description: 市场规模相关类

package types

// ChineseCateringStatistic
// @Description: Gorm模型
type ChineseCateringStatistic struct {
	Id                    uint    `gorm:"primaryKey BIGINT NOT NULL AUTOINCREMENT"`
	Year                  int     `gorm:"INT NULL DEFAULT NULL"`
	TotalAmount           float64 `gorm:"FLOAT NULL DEFAULT NULL"`
	TotalAmountPercentage float64 `gorm:"FLOAT NULL DEFAULT NULL"`
}

// PostAmountAndPercentageByYearFromStatisticRequest
// @Description: 请求订单量和市场占比接口请求类型
type PostAmountAndPercentageByYearFromStatisticRequest struct {
	BeginYear int `json:"beginYear"`
	EndYear   int `json:"endYear"`
}

// PostAmountAndPercentageByYearFromStatisticData
// @Description: 请求订单量和市场占比接口传递的基础数据
type PostAmountAndPercentageByYearFromStatisticData struct {
	Year       string `json:"year"`
	Amount     string `json:"amount"`
	Percentage string `json:"percentage"`
}

// PostAmountAndPercentageByYearFromStatisticResponse
// @Description: 请求订单量和市场占比接口返回内容
type PostAmountAndPercentageByYearFromStatisticResponse struct {
	Code ErrNo                                            `json:"code"`
	Data []PostAmountAndPercentageByYearFromStatisticData `json:"data"`
}
