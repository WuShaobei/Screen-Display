package types

type ChineseCateringBrandStatistic struct {
	Id    uint   `gorm:"primaryKey BIGINT NOT NULL AUTOINCREMENT"`
	Brand string `gorm:"VARCHAR(255) DEFAULT NULL"`
	Price int    `gorm:"INT NULL DEFAULT NULL"`
	Other string `gorm:"VARCHAR(255) NULL DEFAULT NULL"`
}

type SelectChineseCateringBrandStatisticResponse struct {
	Code ErrNo
	Data []string
}
