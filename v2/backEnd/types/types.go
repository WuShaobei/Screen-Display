/*
 * @Author: WuShaobei
 * @Date: 2022-10-09 21:28:14
 * @LastEditTime: 2022-10-09 21:28:27
 * @FilePath: /Screen-Display/v2/backEnd/types/types.go
 * @Description: 基础常量
 */

package types

type IdentityTypes string

const (
	Admin        string = "0" // 管理员
	Investor     string = "1" // 投资人
	Practitioner string = "2" // 从业者
	Tourist      string = "3" // 游客
)

type ErrNo int

const (
	OK             ErrNo = 0
	UserNotExist   ErrNo = 1
	UserHasExisted ErrNo = 1   // 该 Username 已存在
	WrongPassword  ErrNo = 2   // 密码错误
	UnknownError   ErrNo = 255 // 未知错误
)
