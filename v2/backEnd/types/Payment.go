// Package types
// @Description: 薪资相关类

package types

// ChineseCateringPayment
// @Description: Payment Gorm模型
type ChineseCateringPayment struct {
	Id     uint   `gorm:"primaryKey BIGINT NOT NULL AUTOINCREMENT"`
	Year   int    `gorm:"INT NULL DEFAULT NULL"`
	Name   string `gorm:"VARCHAR(255) NULL DEFAULT NULL"`
	Salary int    `gorm:"INT NULL DEFAULT NULL"`
}

// PostAvgSalaryAndCountsByYearFromPaymentRequest
// @Description: 请求平均薪资和就业人数请求类型
type PostAvgSalaryAndCountsByYearFromPaymentRequest struct {
	BeginYear int `json:"beginYear"`
	EndYear   int `json:"endYear"`
}

// PostAvgSalaryAndCountsByYearFromPaymentData
// @Description: 请求平均薪资和就业人数接口传递的基础数据
type PostAvgSalaryAndCountsByYearFromPaymentData struct {
	Year      string `json:"year"`
	AvgSalary string `json:"avgSalary"`
	Counts    string `json:"counts"`
}

// PostAvgSalaryAndCountsByYearFromPaymentResponse
// @Description: 请求平均薪资和就业人数接口返回内容
type PostAvgSalaryAndCountsByYearFromPaymentResponse struct {
	Code ErrNo                                         `json:"code"`
	Data []PostAvgSalaryAndCountsByYearFromPaymentData `json:"data"`
}
