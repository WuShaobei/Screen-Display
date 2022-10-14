package types

type ChineseCateringStatistic struct {
	Id                    uint    `gorm:"primaryKey BIGINT NOT NULL AUTOINCREMENT"`
	Year                  int     `gorm:"INT NULL DEFAULT NULL"`
	TotalAmount           float64 `gorm:"FLOAT NULL DEFAULT NULL"`
	TotalAmountPercentage float64 `gorm:"FLOAT NULL DEFAULT NULL"`
}

type PostAmountAndPercentageByYearFromStatisticRequest struct {
	BeginYear int `json:"beginYear"`
	EndYear   int `json:"endYear"`
}

type PostAmountAndPercentageByYearFromStatisticData struct {
	Year       string `json:"year"`
	Amount     string `json:"amount"`
	Percentage string `json:"percentage"`
}

type PostAmountAndPercentageByYearFromStatisticResponse struct {
	Code ErrNo                                            `json:"code"`
	Data []PostAmountAndPercentageByYearFromStatisticData `json:"data"`
}
