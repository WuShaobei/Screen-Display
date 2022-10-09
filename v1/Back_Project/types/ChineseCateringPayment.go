package types

type ChineseCateringPayment struct {
	Id     uint   `gorm:"primaryKey BIGINT NOT NULL AUTOINCREMENT"`
	Year   int    `gorm:"INT NULL DEFAULT NULL"`
	Name   string `gorm:"VARCHAR(255) NULL DEFAULT NULL"`
	Salary int    `gorm:"INT NULL DEFAULT NULL"`
}

type SelectGetChineseCateringPaymentRequest struct {
	BeginYear int
	EndYear   int
}

type SelectGetChineseCateringPaymentData struct {
	Year        int
	AvgSalary   float64
	CountSalary string
}

type SelectGetChineseCateringPaymentResponse struct {
	Code ErrNo
	Data []SelectGetChineseCateringPaymentData
}
