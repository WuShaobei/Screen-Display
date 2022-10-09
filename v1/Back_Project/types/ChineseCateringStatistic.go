package types

type ChineseCateringStatistic struct {
	Id                    uint    `gorm:"primaryKey BIGINT NOT NULL AUTOINCREMENT"`
	Year                  int     `gorm:"INT NULL DEFAULT NULL"`
	TotalAmount           float64 `gorm:"FLOAT NULL DEFAULT NULL"`
	TotalAmountPercentage float64 `gorm:"FLOAT NULL DEFAULT NULL"`
}

type SelectChineseCateringStatisticRequest struct {
	BeginYear int
	EndYear   int
}

type SelectChineseCateringStatisticData struct {
	Year                  int
	TotalAmount           string
	TotalAmountPercentage string
}

type SelectChineseCateringStatisticResponse struct {
	Code ErrNo
	Data []SelectChineseCateringStatisticData
}
