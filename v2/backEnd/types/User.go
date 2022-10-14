// Package types
// @Description:用户相关类

package types

// ChineseCateringUser
// @Description: User类 Gorm 模型
type ChineseCateringUser struct {
	Id       uint          `gorm:"primaryKey;type:bigint UNSIGNED not null AUTO_INCREMENT"`
	Username string        `gorm:"VARCHAR(255) NULL DEFAULT NULL"`
	Password string        `gorm:"VARCHAR(255) NULL DEFAULT NULL"`
	RealName string        `gorm:"VARCHAR(255) NULL DEFAULT NULL"`
	Identity IdentityTypes `gorm:"enum(0,1,2,3) NULL DEFAULT NULL"`
}

// WhoAmIRequest
// @Description: 查看用户信息接口请求类型
type WhoAmIRequest struct {
	Id string `form:"id" json:"id"`
}

// WhoAmIData
// @Description: 查看用户信息接口传递的基础数据
type WhoAmIData struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	RealName string `json:"realName"`
	Identity string `json:"identity"`
}

// WhoAmIResponse
// @Description: 查看用户信息接口返回内容
type WhoAmIResponse struct {
	Code ErrNo      `json:"code"`
	Data WhoAmIData `json:"data"`
}

// LoginByPasswordRequest
// @Description: 通过用户名和密码登录接口请求类型
type LoginByPasswordRequest struct {
	Username  string `form:"username" json:"username"`
	Password  string `form:"password" json:"password"`
	SetCookie bool   `form:"setCookie" json:"setCookie"`
}

// LoginBySessionKeyRequest
// @Description: 通过 session 值登录接口请求类型
type LoginBySessionKeyRequest struct {
	SessionKey string `form:"sessionKey" json:"sessionKey"`
}

// LoginData
// @Description: 登录接口传递的基础数据
type LoginData struct {
	Id         string `json:"id"`
	SessionKey string `json:"sessionKey"`
}

// LoginResponse
// @Description: 登录接口返回的内容
type LoginResponse struct {
	Code ErrNo     `json:"code"`
	Data LoginData `json:"data"`
}

// LogoutRequest
// @Description: 登出接口请求类型
type LogoutRequest struct {
	SessionKey string `form:"sessionKey" json:"sessionKey"`
}

// LogoutResponse
// @Description: 登出接口返回内容
type LogoutResponse struct {
	Code ErrNo `json:"code"`
}

// RegisterRequest
// @Description: 用户注册接口请求类型
type RegisterRequest struct {
	Username string        `form:"username" json:"username"`
	Password string        `form:"password" json:"password"`
	RealName string        `form:"realName" json:"realName"`
	Identity IdentityTypes `form:"identity" json:"identity"`
}

// RegisterResponse
// @Description: 用户注册接口返回内容
type RegisterResponse struct {
	Code ErrNo `json:"code"`
	Data struct {
		Id string `json:"id"`
	} `json:"data"`
}
