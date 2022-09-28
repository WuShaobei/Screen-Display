package types

type ChineseCateringUser struct {
	Id       uint   `gorm:"primaryKey;type:bigint UNSIGNED not null AUTO_INCREMENT"`
	Username string `gorm:"VARCHAR(255) NULL DEFAULT NULL"`
	Password string `gorm:"VARCHAR(255) NULL DEFAULT NULL"`
	RealName string `gorm:"VARCHAR(255) NULL DEFAULT NULL"`
	Identity string `gorm:"enum(0,1,2,3) NULL DEFAULT NULL"`
}

type LoginRequest struct {
	Username   string
	Password   string
	SessionKey string
}

type LoginResponse struct {
	Code ErrNo
	Data struct {
		Session string
		Id      uint
	}
}

type CreateMemberRequest struct {
	Username string
	Password string
	RealName string
	Identity string
}

type CreateMemberResponse struct {
	Code ErrNo
	Data struct {
		Id uint
	}
}

type WhoAmIRequest struct {
	Id string
}

type WhoAmIResponse struct {
	Code ErrNo
	Data struct {
		Username string
		RealName string
		Identity string
	}
}
