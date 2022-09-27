package types

// ChineseCateringPayment
// Table4
type ChineseCateringPayment struct {
	Id     string `gorm:"primaryKey BIGINT NOT NULL AUTOINCREMENT"`
	Year   int    `gorm:"INT NULL DEFAULT NULL"`
	Name   string `gorm:"VARCHAR(255) NULL DEFAULT NULL"`
	Salary int    `gorm:"INT NULL DEFAULT NULL"`
}

type Api3Request struct {
	BeginYear int
	EndYear   int
}

type Api3Data struct {
	Year        int
	AvgSalary   string
	CountSalary string
}

type Api3Response struct {
	Code ErrNo
	Data []Api3Data
}
