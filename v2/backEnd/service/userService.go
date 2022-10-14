// Package service
// @Description: 用户类接口对应的 Service 层
package service

import (
	"backEnd/manager"

	"backEnd/types"
	"log"

	"github.com/gin-gonic/gin"
)

// WhoAmIService
//
//	@Description: 获取用户信息接口
//	@param c
func WhoAmIService(c *gin.Context) {
	var request types.WhoAmIRequest
	if err := c.ShouldBind(&request); err != nil {
		log.Println(err)
		return
	}

	var userManage manager.UserManager
	data, errNo := userManage.WhoAmI(request.Id)

	if data != (types.WhoAmIData{}) {
		data.Identity = types.GetIdentity(types.IdentityTypes(data.Identity))
		c.JSON(200, types.WhoAmIResponse{Code: errNo, Data: data})
	} else {
		c.JSON(200, types.WhoAmIResponse{Code: types.UserNotExist})
	}
}

// LoginByPasswordService
//
//	@Description: 通过用户名和密码登录接口
//	@param c
func LoginByPasswordService(c *gin.Context) {
	var request types.LoginByPasswordRequest
	if err := c.ShouldBind(&request); err != nil {
		log.Println(err)
		return
	}

	request.Password = types.MD5(request.Password)
	userManager := manager.UserManager{}
	data, errNo := userManager.LoginByPassword(request.Username, request.Password, request.SetCookie)

	if errNo == types.OK {
		c.JSON(200, types.LoginResponse{Code: types.OK, Data: data})
	} else {
		c.JSON(200, types.LoginResponse{Code: errNo})
	}

}

// LoginBySessionKeyService
//
//	@Description: 通过 session 值登录接口
//	@param c
func LoginBySessionKeyService(c *gin.Context) {
	var request types.LoginBySessionKeyRequest
	if err := c.ShouldBind(&request); err != nil {
		log.Println(err)
		return
	}

	userManager := manager.UserManager{}
	data, errNo := userManager.LoginBySessionKey(request.SessionKey)
	if errNo == types.OK {
		c.JSON(200, types.LoginResponse{
			Code: errNo,
			Data: data,
		})
	} else {
		c.JSON(200, types.LoginResponse{
			Code: errNo,
		})
	}
}

// LogoutService
//
//	@Description: 登出接口
//	@param c
func LogoutService(c *gin.Context) {
	var request types.LogoutRequest
	if err := c.ShouldBind(&request); err != nil {
		log.Println(err)
		return
	}

	userManager := manager.UserManager{}
	errNo := userManager.Logout(request.SessionKey)

	c.JSON(200, types.LogoutResponse{Code: errNo})
}

// RegisterService
//
//	@Description: 用户注册接口
//	@param c
func RegisterService(c *gin.Context) {
	var request types.RegisterRequest
	if err := c.ShouldBind(&request); err != nil {
		log.Println(err)
		return
	}

	request.Password = types.MD5(request.Password)
	userManager := manager.UserManager{}
	data, errNo := userManager.Register(request)

	if errNo == types.OK {
		c.JSON(200,
			types.RegisterResponse{
				Code: types.OK,
				Data: struct {
					Id string `json:"id"`
				}{Id: data},
			},
		)
	} else {
		c.JSON(200, types.RegisterResponse{Code: errNo})
	}
}
