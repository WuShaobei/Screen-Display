// Package types
// @Description: 融资情况相关类

package types

// ChineseCateringFundingStatistic
// @Description: 融资情况Gorm模型
type ChineseCateringFundingStatistic struct {
	Id           uint   `gorm:"primaryKey BIGINT NOT NULL AUTOINCREMENT"`
	Brand        string `gorm:"VARCHAR(255) NULL DEFAULT NULL"`
	FundingYear  int    `gorm:"INT NULL DEFAULT NULL"`
	FundingMonth int    `gorm:"INT NULL DEFAULT NULL"`
	FundingRound string `gorm:"VARCHAR(255) NULL DEFAULT NULL"`
	Investor     string `gorm:"VARCHAR(255) NULL DEFAULT NULL"`
}

// GetAllDataFromFundingStatisticData
// @Description: 获取融资情况的所有数据接口传递的基础数据
type GetAllDataFromFundingStatisticData struct {
	Brand        string `json:"brand"`
	Time         string `json:"time"`
	FundingRound string `json:"fundingRound"`
	Investor     string `json:"investor"`
}

// PostAllDataByYearAndMonthFromFundingStatisticResponse
// @Description: 获取融资情况的所有数据接口的返回内容
type PostAllDataByYearAndMonthFromFundingStatisticResponse struct {
	Code ErrNo                                `json:"code"`
	Data []GetAllDataFromFundingStatisticData `json:"data"`
}
