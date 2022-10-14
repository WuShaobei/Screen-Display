package types

type ChineseCateringPayment struct {
	Id     uint   `gorm:"primaryKey BIGINT NOT NULL AUTOINCREMENT"`
	Year   int    `gorm:"INT NULL DEFAULT NULL"`
	Name   string `gorm:"VARCHAR(255) NULL DEFAULT NULL"`
	Salary int    `gorm:"INT NULL DEFAULT NULL"`
}

type PostAvgSalaryAndCountsByYearFromPaymentRequest struct {
	BeginYear int `json:"beginYear"`
	EndYear   int `json:"endYear"`
}

type PostAvgSalaryAndCountsByYearFromPaymentData struct {
	Year      string `json:"year"`
	AvgSalary string `json:"avgSalary"`
	Counts    string `json:"counts"`
}

type PostAvgSalaryAndCountsByYearFromPaymentResponse struct {
	Code ErrNo                                         `json:"code"`
	Data []PostAvgSalaryAndCountsByYearFromPaymentData `json:"data"`
}
