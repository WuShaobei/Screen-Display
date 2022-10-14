package types

type ChineseCateringFundingStatistic struct {
	Id           uint   `gorm:"primaryKey BIGINT NOT NULL AUTOINCREMENT"`
	Brand        string `gorm:"VARCHAR(255) NULL DEFAULT NULL"`
	FundingYear  int    `gorm:"INT NULL DEFAULT NULL"`
	FundingMonth int    `gorm:"INT NULL DEFAULT NULL"`
	FundingRound string `gorm:"VARCHAR(255) NULL DEFAULT NULL"`
	Investor     string `gorm:"VARCHAR(255) NULL DEFAULT NULL"`
}

type GetAllDataFromFundingStatisticData struct {
	Brand        string `json:"brand"`
	Time         string `json:"time"`
	FundingRound string `json:"fundingRound"`
	Investor     string `json:"investor"`
}

type PostAllDataByYearAndMonthFromFundingStatisticResponse struct {
	Code ErrNo                                `json:"code"`
	Data []GetAllDataFromFundingStatisticData `json:"data"`
}
