package types

// ChineseCateringBrandStatistic
// Table1
type ChineseCateringBrandStatistic struct {
	Id    string `gorm:"primaryKey BIGINT NOT NULL AUTOINCREMENT"`
	Brand string `gorm:"INT NULL DEFAULT NULL"`
	Other string `gorm:"VARCHAR(255) NULL DEFAULT NULL"`
}

type Api4Response struct {
	Code ErrNo
	Data []string
}
