package types

type ChineseCateringOnlineOrderStatistic struct {
	Id          uint `gorm:"primaryKey BIGINT NOT NULL AUTOINCREMENT"`
	Year        int  `gorm:"INT NULL DEFAULT NULL"`
	Month       int  `gorm:"INT NULL DEFAULT NULL"`
	OrderAmount int  `gorm:"INT NULL DEFAULT NULL"`
}

type SelectChineseCateringOnlineOrderStatisticRequest struct {
	BeginYear int
	EndYear   int
}

type SelectChineseCateringOnlineOrderStatisticData struct {
	Time        string
	OrderAmount int
}

type SelectChineseCateringOnlineOrderStatisticResponse struct {
	Code           ErrNo
	AllTime        []string
	AllOrderAmount []int
}
