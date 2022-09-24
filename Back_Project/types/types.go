package types

type IdentityTypes int
type ErrNo int

const (
	Root         int = 0 // 管理员
	Investor     int = 1 // 投资人
	Practitioner int = 2 // 从业者
	Tourist      int = 3 // 游客
)

const (
	OK             ErrNo = 0
	ParamInvalid   ErrNo = 1   // 参数不合法
	UserHasExisted ErrNo = 2   // 该 Username 已存在
	UserHasDeleted ErrNo = 3   // 用户已删除
	UserNotExisted ErrNo = 4   // 用户不存在
	WrongPassword  ErrNo = 5   // 密码错误
	LoginRequired  ErrNo = 6   // 用户未登录
	PermDenied     ErrNo = 10  // 没有操作权限
	UnknownError   ErrNo = 255 // 未知错误
)

const (
	Table1ChineseCateringBrandStatistic       int = 1
	Table2ChineseCateringFundingStatistic         = 2
	Table3ChineseCateringOnlineOrderStatistic     = 3
	Table4ChineseCateringPayment                  = 4
	Table5ChineseCateringStatistic                = 5
	Table6ChineseCateringUser                     = 6
)
