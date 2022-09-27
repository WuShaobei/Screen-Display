package types

// ChineseCateringOnlineOrderStatistic
// Table3
type ChineseCateringOnlineOrderStatistic struct {
	Id          string `gorm:"primaryKey BIGINT NOT NULL AUTOINCREMENT"`
	Year        int    `gorm:"INT NULL DEFAULT NULL"`
	Month       int    `gorm:"INT NULL DEFAULT NULL"`
	OrderAmount int    `gorm:"INT NULL DEFAULT NULL"`
}

type Api2Request struct {
	BeginYear int
	EndYear   int
}

type Api2Data struct {
	Time        string
	OrderAmount int
}

type Api2Response struct {
	Code           ErrNo
	AllTime        []string
	AllOrderAmount []int
}
