package types

type ChineseCateringOnlineOrderStatistic struct {
	Id          uint `gorm:"primaryKey BIGINT NOT NULL AUTOINCREMENT"`
	Year        int  `gorm:"INT NULL DEFAULT NULL"`
	Month       int  `gorm:"INT NULL DEFAULT NULL"`
	OrderAmount int  `gorm:"INT NULL DEFAULT NULL"`
}

type PostAmountByYearAndMonthFromOnlineOrderStatisticRequest struct {
	BeginYear int `json:"beginYear"`
	EndYear   int `json:"endYear"`
}

type PostAmountByYearAndMonthFromOnlineOrderStatisticResponse struct {
	Code ErrNo `json:"code"`
	Data struct {
		AllTime   []string `json:"allTime"`
		AllAmount []string `json:"allAmount"`
	} `json:"data"`
}
