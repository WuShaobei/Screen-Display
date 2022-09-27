package types

type IdentityTypes int
type ErrNo int

const (
	Admin        string = "0" // 管理员
	Investor     string = "1" // 投资人
	Practitioner string = "2" // 从业者
	Tourist      string = "3" // 游客
)

const (
	OK             ErrNo = 0
	UserHasExisted ErrNo = 1   // 该 Username 已存在
	WrongPassword  ErrNo = 2   // 密码错误
	UnknownError   ErrNo = 255 // 未知错误
)
