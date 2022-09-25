package types

type LoginRequest struct {
	Username   string
	Password   string
	SessionKey string
}

// 登录成功后需要 Set-Cookie("camp-session", ${value})

type LoginResponse struct {
	// 密码错误范围密码错误状态码
	Code    ErrNo
	Session string
	Data    ChineseCateringUser
}

//
//// 登出
//
//type LogoutRequest struct{}
//
//// 登出成功需要删除 Cookie
//
//type LogoutResponse struct {
//	Code ErrNo
//}
//
//// WhoAmI 接口，用来测试是否登录成功，只有此接口需要带上 Cookie
//
//type WhoAmIRequest struct {
//}
//
//// 用户未登录请返回用户未登录状态码
//
//type WhoAmIResponse struct {
//	Code ErrNo
//	Data ChineseCateringUser
//}
//
//type CreateMemberRequest struct {
//	Username string
//	Password string
//	RealName string
//	Identity IdentityTypes
//}
//
//type CreateMemberResponse struct {
//	Code ErrNo
//	Data struct {
//		UserID string // int64 范围
//	}
//}
//
//type GetMemberRequest struct {
//	Id string
//}
//
//// 如果用户已删除请返回已删除状态码，不存在请返回不存在状态码
//
//type GetMemberResponse struct {
//	Code ErrNo
//	Data ChineseCateringUser
//}
