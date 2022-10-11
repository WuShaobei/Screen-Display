/*
 * @Author: WuShaobei
 * @Date: 2022-10-09 21:02:22
 * @LastEditTime: 2022-10-10 08:47:59
 * @FilePath: /Screen-Display/v2/backEnd/types/User.go
 * @Description:
 */

package types

// ChineseCateringUser
/**
 * @description: 访问数据库基本单位
 */
type ChineseCateringUser struct {
	Id       uint          `gorm:"primaryKey;type:bigint UNSIGNED not null AUTO_INCREMENT"`
	Username string        `gorm:"VARCHAR(255) NULL DEFAULT NULL"`
	Password string        `gorm:"VARCHAR(255) NULL DEFAULT NULL"`
	RealName string        `gorm:"VARCHAR(255) NULL DEFAULT NULL"`
	Identity IdentityTypes `gorm:"enum(0,1,2,3) NULL DEFAULT NULL"`
}

// WhoAmIRequest
/**
 * @description: whoAmI 请求体
 */
type WhoAmIRequest struct {
	Id string `form:"id" json:"id"`
}

// WhoAmIData
/**
 * @description: whoAmI 数据类
 */
type WhoAmIData struct {
	Id       string        `json:"id"`
	Username string        `json:"username"`
	RealName string        `json:"realName"`
	Identity IdentityTypes `json:"identity"`
}

// WhoAmIResponse
/**
 * @description: whoAmI 返回体
 */
type WhoAmIResponse struct {
	Code ErrNo      `json:"code"`
	Data WhoAmIData `json:"data"`
}

type LoginByPasswordRequest struct {
	Username  string `form:"username" json:"username"`
	Password  string `form:"password" json:"password"`
	SetCookie bool   `form:"setCookie" json:"setCookie"`
}

type LoginBySessionKeyRequest struct {
	SessionKey string `form:"sessionKey" json:"sessionKey"`
}

type LoginData struct {
	Id         string `json:"id"`
	SessionKey string `json:"sessionKey"`
}

type LoginResponse struct {
	Code ErrNo     `json:"code"`
	Data LoginData `json:"data"`
}

type LogoutRequest struct {
	SessionKey string `form:"sessionKey" json:"sessionKey"`
}

type LogoutResponse struct {
	Code ErrNo `json:"code"`
}

type RegisterRequest struct {
	Username string        `form:"username" json:"username"`
	Password string        `form:"password" json:"password"`
	RealName string        `form:"realName" json:"realName"`
	Identity IdentityTypes `form:"identity" json:"identity"`
}

type RegisterResponse struct {
	Code ErrNo `json:"code"`
	Data struct {
		Id string `json:"id"`
	} `json:"data"`
}
