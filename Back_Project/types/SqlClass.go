package types

// ChineseCateringBrandStatistic
// Table1
type ChineseCateringBrandStatistic struct {
	Id    string `gorm:"primaryKey BIGINT NOT NULL AUTOINCREMENT"`
	Brand string `gorm:"INT NULL DEFAULT NULL"`
	Other string `gorm:"VARCHAR(255) NULL DEFAULT NULL"`
}

// ChineseCateringFundingStatistic
// Table2
type ChineseCateringFundingStatistic struct {
	Id           string `gorm:"primaryKey BIGINT NOT NULL AUTOINCREMENT"`
	Brand        string `gorm:"VARCHAR(255) NULL DEFAULT NULL"`
	FundingYear  int    `gorm:"INT NULL DEFAULT NULL"`
	FundingMonth int    `gorm:"INT NULL DEFAULT NULL"`
	FundingRound string `gorm:"VARCHAR(255) NULL DEFAULT NULL"`
	Investor     string `gorm:"VARCHAR(255) NULL DEFAULT NULL"`
}

// ChineseCateringOnlineOrderStatistic
// Table3
type ChineseCateringOnlineOrderStatistic struct {
	Id          string `gorm:"primaryKey BIGINT NOT NULL AUTOINCREMENT"`
	Year        int    `gorm:"INT NULL DEFAULT NULL"`
	Month       int    `gorm:"INT NULL DEFAULT NULL"`
	OrderAmount int    `gorm:"INT NULL DEFAULT NULL"`
}

// ChineseCateringPayment
// Table4
type ChineseCateringPayment struct {
	Id     string `gorm:"primaryKey BIGINT NOT NULL AUTOINCREMENT"`
	Year   int    `gorm:"INT NULL DEFAULT NULL"`
	Name   string `gorm:"VARCHAR(255) NULL DEFAULT NULL"`
	Salary int    `gorm:"INT NULL DEFAULT NULL"`
}

// ChineseCateringStatistic
// Table5
type ChineseCateringStatistic struct {
	Id string `gorm:"primaryKey BIGINT NOT NULL AUTOINCREMENT"`

	Year                  int     `gorm:"INT NULL DEFAULT NULL"`
	TotalAmount           float64 `gorm:"FLOAT NULL DEFAULT NULL"`
	TotalAmountPercentage float64 `gorm:"FLOAT NULL DEFAULT NULL"`
}

// ChineseCateringUser
// Table6
type ChineseCateringUser struct {
	Id       string        `gorm:"primaryKey BIGINT NOT NULL AUTOINCREMENT"`
	Username string        `gorm:"VARCHAR(255) NULL DEFAULT NULL"`
	Password string        `gorm:"VARCHAR(255) NULL DEFAULT NULL"`
	RealName string        `gorm:"VARCHAR(255) NULL DEFAULT NULL"`
	Identity IdentityTypes `gorm:"enum(0,1,2,3) NULL DEFAULT NULL"`
}
