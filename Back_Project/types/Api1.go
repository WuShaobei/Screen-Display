package types

// ChineseCateringStatistic
// Table5
type ChineseCateringStatistic struct {
	Id                    uint    `gorm:"primaryKey BIGINT NOT NULL AUTOINCREMENT"`
	Year                  int     `gorm:"INT NULL DEFAULT NULL"`
	TotalAmount           float64 `gorm:"FLOAT NULL DEFAULT NULL"`
	TotalAmountPercentage float64 `gorm:"FLOAT NULL DEFAULT NULL"`
}

type Api1Request struct {
	BeginYear int
	EndYear   int
}

type Api1Data struct {
	Year                  int
	TotalAmount           string
	TotalAmountPercentage string
}

type Api1Response struct {
	Code ErrNo
	Data []Api1Data
}
