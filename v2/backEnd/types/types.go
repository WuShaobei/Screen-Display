// Package types
// @Description: 类型定义及常量设置

package types

// IdentityTypes
// 用户类型
type IdentityTypes string

const (
	Admin        IdentityTypes = "0" // 管理员
	Investor     IdentityTypes = "1" // 投资人
	Practitioner IdentityTypes = "2" // 从业者
	Tourist      IdentityTypes = "3" // 游客
)

// ErrNo
// 异常值返回
type ErrNo int

const (
	OK             ErrNo = 0
	UserNotExist   ErrNo = 1
	UserHasExisted ErrNo = 2   // 该 Username 已存在
	WrongPassword  ErrNo = 3   // 密码错误
	UnknownError   ErrNo = 255 // 未知错误
)

// 数据表有效时间
const (
	StatisticMinYear int = 2014
	StatisticMaxYear int = 2021
	PaymentMinYear   int = 2018
	PaymentMaxYear   int = 2021
	OrderMinYear     int = 2010
	OrderMaxYear     int = 2021
)
