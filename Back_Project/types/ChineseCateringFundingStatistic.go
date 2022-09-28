package types

type ChineseCateringFundingStatistic struct {
	Id           uint   `gorm:"primaryKey BIGINT NOT NULL AUTOINCREMENT"`
	Brand        string `gorm:"VARCHAR(255) NULL DEFAULT NULL"`
	FundingYear  int    `gorm:"INT NULL DEFAULT NULL"`
	FundingMonth int    `gorm:"INT NULL DEFAULT NULL"`
	FundingRound string `gorm:"VARCHAR(255) NULL DEFAULT NULL"`
	Investor     string `gorm:"VARCHAR(255) NULL DEFAULT NULL"`
}

type SelectChineseCateringFundingStatisticResponse struct {
	Code ErrNo
	Data []struct {
		Time         string
		Brand        string
		FundingRound string
		Investor     string
	}
}
